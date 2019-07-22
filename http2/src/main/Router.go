package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/template/").Handler(http.StripPrefix("/template/",http.FileServer(http.Dir("template/"))))
	//common.PathPrefix("/js/").Handler(http.StripPrefix("/js/",http.FileServer(http.Dir("template/"))))

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
