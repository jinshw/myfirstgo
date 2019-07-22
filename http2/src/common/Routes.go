package common

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	//Route{Name: "Index", Method: "GET", Pattern: "/", HandlerFunc: Index},
	Route{Name: "TodoIndex", Method: "GET", Pattern: "/*", HandlerFunc: TodoIndex},
	Route{Name: "TodoShow", Method: "GET", Pattern: "/todos/{todoId}", HandlerFunc: TodoShow},
}