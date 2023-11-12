package main

import (
	"golang-form-value/handler"
	"log"
	"net/http"
)

var port = ":8080"

func main() {
	http.HandleFunc("/", handler.HandleRoot)
	http.HandleFunc("/process", handler.HandleProcess)

	log.Printf("server running on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
