package handlers

import (
	"log/slog"
	"net/http"
)

func handleHandlerError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		slog.Error(err.Error())
	}
}
