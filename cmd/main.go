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
	dbUrl string
}

func main() {
	cfg := config{
		dbUrl: "postgres://postgres:postgres@localhost:5432/portal_back?sslmode=disable",
		addr:  ":8080",
	}

	db, err := connectToDatabase(cfg.dbUrl)
	handleCriticalError(err)

	err = service.SetupServiceLayer(db)
	handleCriticalError(err)

	err = startHTTPServer(cfg.addr)
	slog.Error(err.Error())
}

func getConfigFromEnvironment() (*config, error) {
	var cfg config

	cfg.dbUrl = os.Getenv("DATABASE_URL")
	if cfg.dbUrl == "" {
		return nil, errors.New("environment variable \"DATABASE_URL\" not set")
	}

	cfg.addr = os.Getenv("ADDR")
	if cfg.dbUrl == "" {
		return nil, errors.New("environment variable \"ADDR\" not set")
	}

	return &cfg, nil
}

func connectToDatabase(url string) (*sql.DB, error) {
	slog.Info("Connecting to the database")

	conn, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func startHTTPServer(addr string) error {
	s := &http.Server{
		Addr:    addr,
		Handler: handlers.SetupRoutes(),
	}

	slog.Info("Starting web server", "addr", ":8080")
	return s.ListenAndServe()
}

func handleCriticalError(err error) {
	if err == nil {
		return
	}

	slog.Error("Something wrong is not right", "error", err.Error())
	os.Exit(-1)
}
