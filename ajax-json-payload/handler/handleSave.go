package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleSave(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Only accept POST method request", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(req.Body)
	payload := struct {
		Age    int    `json:"age"`
		Name   string `json:"name"`
		Gender string `json:"gender"`
	}{}

	if err := decoder.Decode(&payload); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Println("Error Bro")
		return
	}

	message := fmt.Sprintf(
		"Hello, my name is %s. I'm %d year old, %s",
		payload.Name,
		payload.Age,
		payload.Gender,
	)
	res.Write([]byte(message))
}
