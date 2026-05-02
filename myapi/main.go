package main

import (
	"beginner/myapi/handlers"
	"beginner/myapi/middlewares"
	"fmt"
	"net/http"
)

// Create a struct
// http.ServerMux() must be private and must be a pointer
// implement http.handler
// add get, put , post , patch, delete
// get type func string, router handler // error return type

// type AppHandler func(w http.ResponseWriter, r *http.Request) error
// app := NewApp()
// app.Get('/',func(w httpResponseWiter, r *httpRequest)error{

// })

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

	mux.HandleFunc("/login",
		middlewares.Logger(
			middlewares.MethodCheck(http.MethodPost,
				handlers.Login),
		))

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Server failed:", err)
	}
}
