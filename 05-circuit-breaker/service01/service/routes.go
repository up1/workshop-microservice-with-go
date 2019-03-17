package service

import "net/http"

// Defines a single route, e.g. a human readable name, HTTP method, pattern the function that will execute when the route is called.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// Initial routes
var routes = Routes{

	Route{
		"CallService 02", // Name
		"GET",            // HTTP method
		"/call/{state}",  // Route pattern
		CallService02,
	},
}
