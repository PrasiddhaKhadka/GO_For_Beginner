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

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/user", getUser)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		http.Error(nil, "Server failed", http.StatusInternalServerError)
	}

}
