package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/stef-k/socketizer-service/utils"
	"github.com/justinas/alice"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		//handler = utils.Logger(handler, route.Name)

		// middlware handlers
		handlers := alice.New(utils.LoggingHandler, utils.RecoveryHandler)

		router.Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handlers.Then(handler))

	}

	for _, route := range apiv1 {
		var handler http.Handler

		handler = route.HandlerFunc
		//handler = utils.Logger(handler, route.Name)

		// middlware handlers
		handlers := alice.New(utils.LoggingHandler, utils.RecoveryHandler)

		router.Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handlers.Then(handler))

	}

	return router
}
