package handlers

import (
	"fmt"
	login "membros-web/internal/views"
	"net/http"
)

func renderLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = render(w, r, login.LoginPage())
}

func processLoginForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Processing the login form...")
}
