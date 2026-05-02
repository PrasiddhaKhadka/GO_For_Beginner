package main

import "fmt"

type paymenter interface {
pay(amount float32)
// refund(amout float32)
// account()
}

type payment struct {
gateway paymenter
}

// Open close principle
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

type payPal struct {
}

func (p payPal) pay(amount float32) {
fmt.Println("Making payment using paypal", amount)
}

func main() {

    // stripePayment := stripePay{}

    // razorPayment := razorPay{}

    paypalPayment := payPal{}

    newPayment := payment{
    	// gateway: stripePayment,
    	// gateway: razorPayment,
    	gateway: paypalPayment,
    }

    newPayment.makePayment(100)

}
