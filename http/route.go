package main

import (
	"fmt"
	"net/http"
)

func waveGreetings(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed!!", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Hello!!! from top of the world!!")
}

func waveGoodBye(w http.ResponseWriter, r *http.Request) {
	// this takes all HTTP method by default
	fmt.Fprintln(w, "Bye!!! from top of the world!!")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed!!!", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Hello, World!")
}

func main() {

	// Registering the new routing path
	mux := http.NewServeMux()

	mux.HandleFunc("/greeting", waveGreetings)
	mux.HandleFunc("/bye", waveGoodBye)
	mux.HandleFunc("/about", aboutPage)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Server failed:", err)
	}

}
