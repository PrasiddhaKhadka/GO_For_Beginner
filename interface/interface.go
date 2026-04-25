package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorPay struct {
}

func (r razorPay) pay(amount float32) {
	fmt.Println("Making payment using razorpay", amount)
}

type stripePay struct {
}

func (s stripePay) pay(amount float32) {
	fmt.Println("Making payment using stripepay", amount)
}

func main() {

	// stripePayment := stripePay{}

	razorPayment := razorPay{}

	newPayment := payment{
		// gateway: stripePayment,
		gateway: razorPayment,
	}

	newPayment.makePayment(100)

}
