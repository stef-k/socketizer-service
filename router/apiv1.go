package router

import 	"github.com/stef-k/socketizer-service/controllers"

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
		"POST",
		"/service/api/v1/pool-info",
		controllers.PoolInfo,
	},
	// WordPress API
	// refresh client specific post for a given post url
	Route {
		"refresh-post",
		"POST",
		"/service/api/v1/wordpress/cmd/client/refresh/post/",
		controllers.ClientRefreshPost,
	},
} // Routes

