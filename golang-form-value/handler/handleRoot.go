package handler

import (
	"html/template"
	"net/http"
)

func HandleRoot(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		tmpl := template.Must(template.New("form").ParseFiles("views/template.html"))
		err := tmpl.Execute(res, nil)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		return
	default:
		http.Error(res, "", http.StatusBadRequest)
	}
}
