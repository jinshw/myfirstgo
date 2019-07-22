package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {
	port := ":8080"
	router := NewRouter()
	fmt.Println("http service start success!port"+port)
	log.Fatal(http.ListenAndServe(port, router))

}