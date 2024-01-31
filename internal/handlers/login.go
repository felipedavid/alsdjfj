package handlers

import (
	"fmt"
	"net/http"
)

func renderLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = renderPage(w, http.StatusOK, "login.gohtml", nil)
}

func processLoginForm(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Processing the login form...")
}
