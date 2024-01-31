package handlers

import "net/http"

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))
	mux.HandleFunc("GET /login", renderLoginPage)
	mux.HandleFunc("POST /login", processLoginForm)

	return mux
}
