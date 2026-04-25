package main

import "fmt"

// enumerated types

// FIRST WAY 🔥
// type OrderStatus int

// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

// SECOND WAY 📍

type OrderStatus string

const (
	Received OrderStatus = "received"
	//  OTHER TYP INHERITED !!!!!
	Confirmed = "confirmed"
	Prepared  = "prepared"
	Delivered = "delivered"
)

func changeOrderStatus(status OrderStatus) {

	fmt.Println("Order status has been changed: ", status)
}

func main() {

	changeOrderStatus(Received)
}
