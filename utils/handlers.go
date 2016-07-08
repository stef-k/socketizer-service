// Package utils contains the following utilities:
// * LoggingHandler used for logging each handler call
// * RecoveryHandler used to recover from errors
// The above handlers used in combination with
// the package "github.com/justinas/alice" which
// is used to chain them in the router like:
// handlers := alice.New(utils.LoggingHandler, utils.RecoveryHandler)
// The log functionality is provided by the package
// https://github.com/jbrodriguez/mlog a simple rotating logger
package utils

import (
	"net/http"
	"time"
	"github.com/jbrodriguez/mlog"
	"fmt"
)

// LoggingHandler a simple logging http handler
func LoggingHandler(next http.Handler) http.Handler  {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		//log.Printf(
		//	"%s %s %s",
		//	r.Method,
		//	r.RequestURI,
		//	time.Since(start),
		//)
		mlog.Info(fmt.Sprintf("%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),))
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// RecoveryHandler recover from various handler errors
func RecoveryHandler(next http.Handler) http.Handler  {

	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				mlog.Warning(fmt.Sprintf("Panic: %+v" , err))
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
