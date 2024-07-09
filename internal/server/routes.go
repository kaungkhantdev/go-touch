package server

import (
	"encoding/json"
	"net/http"
)

var data []User;

type Response struct {
	Message string `json:"msg"`
	Status  string `json:"status"`
	Data []User `json:"data"`
}

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
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
		Data: data,
	}
	ReusableMethod(response, w)
}


/** CREATE */
func create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
        return
	}

	response := Response{
		Message: "Successfully created",
		Status: "Success",
		Data: []User{user},
	}

	data = append(data, user)


	ReusableMethod(response, w)
}

type GetId struct {
	Id string `json:"id"`
}

/** GET ONE */
func getOne(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { 
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := GetId{}
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil { 
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
        return
	}

	user := User{}
	for _, v := range data { 
		if v.Id == id.Id {
			user = v
		 }
	}

	response := Response{
		Message: "Get one by id",
		Status: "Success",
		Data: []User{user},
	}

	ReusableMethod(response, w)

}

/** GET ALL */
func getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { 
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := Response{
		Message: "Get all data",
		Status: "Success",
		Data: data,
	}

	ReusableMethod(response, w)

}

/** UPDATE */
func update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut { 
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	updateUser := User{}
	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
        return
	}

	for i, v := range data { 
		if v.Id == updateUser.Id {
			
			data[i].Name = updateUser.Name
			data[i].Email = updateUser.Email
			response := Response{
				Message: "Get all data",
				Status: "Success",
				Data: []User{data[i]},
			}
		
			ReusableMethod(response, w)

			return

		}
	}

	http.Error(w, "User is not found", http.StatusNotFound);

}


func RegisterRoutes() {
	http.HandleFunc("/", WelcomeMethod)
	http.HandleFunc("/create", create)
	http.HandleFunc("/get", getOne)
	http.HandleFunc("/get-all", getAll)
	http.HandleFunc("/update", update)
}