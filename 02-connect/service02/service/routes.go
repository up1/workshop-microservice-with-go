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
		"ProductDetail", // Name
		"GET",           // HTTP method
		"/hello",        // Route pattern
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write([]byte("{\"result\":\"OK from service 02\"}"))
		},
	},
}
