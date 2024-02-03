package handlers

import (
	"database/sql"
	"membros-web/internal/middleware"
	"net/http"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
)

var sessionManager *scs.SessionManager

// SetupHandlers sets up the routes and middleware for our application.
// db is required to be able to create a new session manager.
func SetupHandlers(db *sql.DB) http.Handler {
	// Setup the session manager
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Store = postgresstore.New(db)

	mux := http.NewServeMux()

	// Serving our static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))

	// Defining our routes
	mux.HandleFunc("GET /login", renderLoginPage)
	mux.HandleFunc("POST /login", processLoginForm)

	// Setting up our middleware chain
	var handler http.Handler = mux

	handler = sessionManager.LoadAndSave(mux)
	handler = middleware.LogRequest(handler)

	return handler
}
