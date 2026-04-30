package main

import (
	"fmt"
	"time"
)

// Simulates a slow database query
func fetchUserFromDB(userID string) string {
	fmt.Println("DB query started...")
	time.Sleep(5 * time.Second)
	fmt.Println("DB query finished")
	return fmt.Sprintf("%s", userID)
}

// Simulates an HTTP handler
func handleRequest() {
	result := fetchUserFromDB("1")
	fmt.Println("Result:", result)
}

func main() {

	go handleRequest()

	time.Sleep(1 * time.Second)
	fmt.Println("Client disconnected — but DB query is STILL running!")
	time.Sleep(5 * time.Second)

}
