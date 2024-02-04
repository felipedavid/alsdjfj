package handlers

import "net/http"

func renderHomePage(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, http.StatusOK, "home.gohtml", nil)
	handleHandlerError(err, w)
}
