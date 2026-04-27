package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func createUser(w http.ResponseWriter, r *http.Request) {

	// 1. Check method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. Decode the JSON body into our struct
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Created user: %s, email: %s, age: %d\n", req.Name, req.Email, req.Age)

}
func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/user", createUser)

	http.ListenAndServe(":8080", mux)
}
