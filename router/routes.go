package router

import (
	"net/http"
	"projects.iccode.net/stef-k/socketizer-service/controllers"
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
		"/service/live/{host}",
		controllers.Live,
	},
	// static index
	Route{
		"index",
		"GET",
		"/service/",
		controllers.Index,
	},

}
