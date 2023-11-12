package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var port = ":8080"

func main() {
	http.HandleFunc("/", ActionIndex)

	log.Printf("server running on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	data := []struct {
		Age  int
		Name string
	}{
		{Age: 24, Name: "Richard Grayson"},
		{Age: 23, Name: "Jason Tod"},
		{Age: 22, Name: "Tim Drake"},
		{Age: 21, Name: "Damian Wayne"},
	}

	w.Header().Set("Content-Type", "aplication/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
