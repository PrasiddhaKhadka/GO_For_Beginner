package main

import (
	"beginner/myapi/handlers"
	"beginner/myapi/middlewares"
	"beginner/myapi/utils"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/user",
		middlewares.Logger(
			middlewares.Auth(
				middlewares.MethodCheck(http.MethodGet,
					utils.Handler(handlers.GetUser),
				),
			),
		),
	)

	mux.HandleFunc("/login",
		middlewares.Logger(
			middlewares.MethodCheck(http.MethodPost,
				utils.Handler(handlers.Login)),
		))

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Server failed:", err)
	}
}
