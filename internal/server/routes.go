package server

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"msg"`
	Status  string `json:"status"`
	Data []string `json:"data"`
}

func ReusableMethod(response Response, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func WelcomeMethod(w http.ResponseWriter, r *http.Request) {
	response := Response{ 
		Message: "Welcome to the Go server main page",
		Status: "OK",
		Data: []string{"Hello", "World"},
	}
	ReusableMethod(response, w)
}

func TestFun(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Welcome from GO language",
		Status: "Success",
	}

	ReusableMethod(response, w)
}


func RegisterRoutes() {
	http.HandleFunc("/", WelcomeMethod)
	http.HandleFunc("/test", TestFun);
}