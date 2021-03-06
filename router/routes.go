package router

import (
	"net/http"
	"github.com/stef-k/socketizer-service/controllers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	// the websocket endpoing
	Route{
		"live",
		"GET",
		"/service/wordpress/live/{host}",
		controllers.Live,
	},
}
