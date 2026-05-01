// BUFFERED !!
package main

import "fmt"

type Email struct {
	from string
	to   string
	body string
}

func main() {
	fmt.Println("BUFFERED!!")

	//  email sending channel

	// 100 :-> can send around 100 data without (un)-blocking but if it goes above 100 then it will be blocking

	email := make(chan Email, 100)

	for i := 0; i < 100; i++ {
		email <- Email{
			from: "test@gmail.com",
			to:   "xyz@gmail.com",
			body: "Hello! this is a testing email!!!",
		}
	}

	for i := 0; i < 100; i++ {
		result := <-email
		fmt.Println(result)
	}

}
