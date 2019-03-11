package main

import (
	"net/http"
	"time"
	"log"
)

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("CofoxAPI: %s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start))
	})
}
