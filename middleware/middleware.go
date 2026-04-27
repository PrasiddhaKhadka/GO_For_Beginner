package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ---- helper ----
func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// ---- middleware ----
func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)
		next(w, r)
		fmt.Printf("← done\n")
	}
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "mysecrettoken" {
			writeJson(w, http.StatusUnauthorized, map[string]string{
				"error": "unauthorized",
			})
			return
		}
		next(w, r)
	}
}

func methodCheckMiddleware(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			writeJson(w, http.StatusMethodNotAllowed, map[string]string{
				"error": "method not allowed",
			})
			return
		}
		next(w, r)
	}
}

// ---- handlers ----

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUser(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "Ramesh", Email: "ramesh@gmail.com"}
	writeJson(w, http.StatusOK, user)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/user", logger(authMiddleware(methodCheckMiddleware(http.MethodGet, getUser))))

	fmt.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Server failed:", err)
	}

}
