package main

import (
	"fmt"
	"net/http"
)

func waveGreetings(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	// READING HEADERS !!!
	userAgent := r.Header.Get("User-Agent")
	token := r.Header.Get("Authorization")

	// Reading Parameters!!
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	if name == "" {
		name = "stranger"
	}

	fmt.Fprintf(w, "Your browser is: %s\n", userAgent)
	fmt.Fprintf(w, "Your token is: %s\n", token)

	fmt.Fprintf(w, "Hello %v! You are %v years old.\n", name, age)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/wave", waveGreetings)

	http.ListenAndServe(":8080", mux)
}
