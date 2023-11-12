package main

import (
	"ajax-multiple-file-upload/handler"
	"log"
	"net/http"
)

var port = ":8080"

func main() {
	http.HandleFunc("/", handler.HandleIndex)
	http.HandleFunc("/upload", handler.HandleUpload)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	log.Printf("server running on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
