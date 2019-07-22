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
	return router
}
