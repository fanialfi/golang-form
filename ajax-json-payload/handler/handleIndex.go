package handler

import (
	"html/template"
	"net/http"
)

func HandleIndex(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Only accept GET request method", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("views/view.html"))
	if err := tmpl.Execute(res, nil); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
