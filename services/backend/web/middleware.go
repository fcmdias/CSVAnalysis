package web

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// enableCORS wraps an http.Handler with CORS support. By applying this middleware,
// the server will include the appropriate CORS headers in responses. This allows
// cross-origin requests to be made to your server.
//
// Usage:
//
//	http.Handle("/path", enableCORS(yourHandler))
func EnableCORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")             // Allow any origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")    // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Allowed headers

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// PanicRecoveryMiddleware is a middleware that recovers from panics and writes a generic error response.
func PanicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Info("Recovered from panic", "error", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// LoggingMiddleware wraps the provided HTTP handler with logging functionality.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log before handling the request
		start := time.Now()
		slog.Info("stated", "method", r.Method, "path", r.URL.Path)

		customWriter := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(customWriter, r)

		duration := time.Since(start)
		slog.Info("completed", "path", fmt.Sprint(r.URL.Path), "duration", fmt.Sprint(duration), "status code", fmt.Sprint(customWriter.statusCode))
	})
}
