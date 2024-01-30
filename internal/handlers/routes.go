package handlers

import "net/http"

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /login", renderLoginPage)
	mux.HandleFunc("POST /login", processLoginForm)

	return mux
}
