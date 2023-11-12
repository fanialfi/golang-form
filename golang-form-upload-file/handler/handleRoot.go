package handler

import (
	"html/template"
	"net/http"
)

func HandleRoot(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("files/view.html"))
	err := tmpl.Execute(res, nil)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
