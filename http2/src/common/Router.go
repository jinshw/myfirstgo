package common

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(configMap map[string]string) *mux.Router {
	reqURL := configMap["reqURL"]
	root := configMap["root"]
	router := mux.NewRouter().StrictSlash(true)
	if reqURL == "/" || reqURL == "" {
		router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(root))))
	} else {
		router.PathPrefix("/" + reqURL + "/").Handler(http.StripPrefix("/"+reqURL+"/", http.FileServer(http.Dir(root))))
	}
	//router.PathPrefix("/" + reqURL + "/").Handler(http.StripPrefix("/"+reqURL+"/", http.FileServer(http.Dir("D:/tools"))))
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
