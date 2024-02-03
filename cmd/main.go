package main

import (
	"database/sql"
	"errors"
	"log/slog"
	"membros-web/internal/handlers"
	"membros-web/internal/service"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type config struct {
	addr  string
	dbURL string
}

func main() {
	cfg := config{
		dbURL: "postgres://postgres:postgres@127.0.0.1:5432/portal_back?sslmode=disable",
		addr:  "127.0.0.1:8080",
	}

	slog.Info("Connecting to the database")
	db, err := connectToDatabase(cfg.dbURL)
	handleCriticalError(err)

	slog.Info("Ensuring sessions table exists")
	err = ensureSessionsTableExists(db)
	handleCriticalError(err)

	err = service.SetupServiceLayer(db)
	handleCriticalError(err)

	mux := handlers.SetupHandlers(db)

	slog.Info("Starting web server", "addr", ":8080")
	err = startHTTPServer(cfg.addr, mux)
	slog.Error(err.Error())
}

func getConfigFromEnvironment() (*config, error) {
	var cfg config

	cfg.dbURL = os.Getenv("DATABASE_URL")
	if cfg.dbURL == "" {
		return nil, errors.New("environment variable \"DATABASE_URL\" not set")
	}

	cfg.addr = os.Getenv("ADDR")
	if cfg.dbURL == "" {
		return nil, errors.New("environment variable \"ADDR\" not set")
	}

	return &cfg, nil
}

func connectToDatabase(url string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func ensureSessionsTableExists(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS sessions (
			token TEXT PRIMARY KEY,
			data BYTEA NOT NULL,
			expiry TIMESTAMPTZ NOT NULL
		);
		
		CREATE INDEX IF NOT EXISTS sessions_expiry_idx ON sessions (expiry);
	`)

	return err
}

func startHTTPServer(addr string, mux http.Handler) error {
	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return s.ListenAndServe()
}

func handleCriticalError(err error) {
	if err == nil {
		return
	}

	slog.Error("Something wrong is not right", "error", err.Error())
	os.Exit(-1)
}
