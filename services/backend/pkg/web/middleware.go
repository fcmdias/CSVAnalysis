package web

import (
	"net/http"
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
