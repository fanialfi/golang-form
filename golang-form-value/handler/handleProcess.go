package handler

import (
	"html/template"
	"net/http"
)

func HandleProcess(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		tmpl := template.Must(template.New("result").ParseFiles("views/template.html"))

		// ParseForm memparsing form data dari view sebelum akhirnya diambil data datanya
		if err := req.ParseForm(); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		// FormValue digunakan untuk pengambilan data dari view,
		// dimana key nya adalah name dari input form dari view
		name := req.FormValue("name")
		message := req.FormValue("message")

		data := map[string]string{"name": name, "message": message}

		if err := tmpl.Execute(res, data); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	default:
		http.Error(res, "", http.StatusBadRequest)
	}
}
