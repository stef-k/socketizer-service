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
	"errors"
	"runtime"
	"strings"
)

// give line number on panic
func identifyPanic() string {
	var name, file string
	var line int
	var pc [16]uintptr

	n := runtime.Callers(3, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		if !strings.HasPrefix(name, "runtime.") {
			break
		}
	}

	switch {
	case name != "":
		return fmt.Sprintf("%v:%v", name, line)
	case file != "":
		return fmt.Sprintf("%v:%v", file, line)
	}

	return fmt.Sprintf("pc:%x", pc)
}

// LoggingHandler a simple logging http handler
func LoggingHandler(next http.Handler) http.Handler  {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

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
		var err error
		defer func() {
			r := recover()
			if r != nil {
				mlog.Info(identifyPanic())
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknonw error")
				}
				mlog.Warning("recover from %s", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
