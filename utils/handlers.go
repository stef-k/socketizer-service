package utils

import (
	"log"
	"net/http"
	"time"
)

// LoggingHandler a simple logging http handler
func LoggingHandler(next http.Handler) http.Handler  {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// RecoveryHandler recover from various handler errors
func RecoveryHandler(next http.Handler) http.Handler  {

	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
