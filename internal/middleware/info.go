package middleware

import (
	"log/slog"
	"net/http"
)

type statusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

// LogRequest logs the request to the console
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{
			ResponseWriter: w,
			Status:         200,
		}

		next.ServeHTTP(rec, r)

		slog.Info("Incoming Request", "remote_addr", r.RemoteAddr, "proto", r.Proto, "method", r.Method, "url", r.URL.RequestURI(), "status", rec.Status, "status_text", http.StatusText(rec.Status))
	})
}
