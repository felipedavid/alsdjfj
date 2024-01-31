package handlers

import (
	"membros-web/internal/service"
	"net/http"
)

func renderLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = renderPage(w, http.StatusOK, "login.gohtml", nil)
}

func processLoginForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	err := service.Login(email, password)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
