package main

import (
	"ajax-json-payload/handler"
	"log"
	"net/http"
)

var port = ":8080"

func main() {
	http.HandleFunc("/", handler.HandleIndex)
	http.HandleFunc("/save", handler.HandleSave)

	// handle static file
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	log.Printf("Server running on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
