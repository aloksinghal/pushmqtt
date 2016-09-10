package pushapi

import (
	"encoding/json"
	"fmt"
	_ "github.com/gorilla/mux"
	"net/http"
)

type ResponseStruct struct {
	Message string
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	//ServerLogger.Println("thank god")
	fmt.Fprint(w, "Welcome! \n")
}


func PublishMessage(w http.ResponseWriter, r *http.Request) {
	ServerLogger.Println("trying to publish")
	var requestBody map[string]interface{}
	var isValid bool
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	requestBody, isValid = ValidateRequestBody(requestBody)
	if !isValid {
		http.Error(w, "Invalid message details", http.StatusBadRequest)
		return
	}
    result := PublishMessageonMQTT(requestBody)
    response := ResponseStruct{result}
	js, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}