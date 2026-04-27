package main

import (
	"beginner/myapi/handlers"
	"beginner/myapi/middlewares"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/user",
		middlewares.Logger(
			middlewares.Auth(
				middlewares.MethodCheck(http.MethodGet,
					handlers.GetUser,
				),
			),
		),
	)

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Server failed:", err)
	}
}
