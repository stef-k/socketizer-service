package router

import 	"projects.iccode.net/stef-k/socketizer-service/controllers"

type Apiv1 []Route

var apiv1 = Routes{
	// broadcast to Pool
	Route{
		"broadcast-pool",
		"GET",
		"/service/api/v1/broadcast-pool/{msg}",
		controllers.BroadcastPool,
	},
	// broadcast to a specified Domain
	Route{
		"broadcast-domain",
		"GET",
		"/service/api/v1/broadcast-domain/{host}/{msg}",
		controllers.BroadcastDomain,
	},
	// get report
	Route{
		"pool-info",
		"GET",
		"/service/api/v1/pool-info",
		controllers.PoolInfo,
	},
	// WordPress API
	// refresh client specific post for a given post url
	// {postUrl} the permalink of the post
	// {host} to be renamed to key - the unique ID of the subscriber
	Route {
		"refresh-post",
		"POST",
		"/service/api/v1/wordpress/cmd/client/refresh/post/",
		controllers.ClientRefreshPost,
	},
} // Routes
