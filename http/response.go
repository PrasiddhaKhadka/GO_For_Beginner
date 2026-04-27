package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// json Handler Function
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// when u write (w) it directly writes the data goes straight to the client immediately.
	json.NewEncoder(w).Encode(data)
}

// httpHandler
func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return
	}

	user := User{Name: "Ramesh", Email: "ramesh@gmail.com", Age: 25}
	writeJSON(w, http.StatusOK, user)
}

// creating user
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return
	}

	var req User

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusInternalServerError, 0)
		return
	}

	if req.Name == "admin" {
		writeJSON(w, http.StatusOK, map[string]string{"message": "Welcome admin!"})
	} else {
		writeJSON(w, http.StatusOK, map[string]string{"message": "Invalid credentials"})
	}

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/user", getUser)
	mux.HandleFunc("/login", login)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		http.Error(nil, "Server failed", http.StatusInternalServerError)
	}

}
