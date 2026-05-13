// ITS A DEPENDENCY INJECTION USING CONSTRUCTOR!!

package main

import "fmt"

type Payment interface {
	Payment(amount int, from string, to string) error
}

type KhaltiPayment struct {
	secret_key string
	public_key string
	live_key   string
}

func (k *KhaltiPayment) Payment(amount int, from string, to string) error {
	fmt.Println("***** Initiating from KhalitPayment *****", amount, from, to)
	return nil
}

type EsewaPayment struct {
	merchant_id string
	secret_key  string
}

func (e *EsewaPayment) Payment(amount int, from string, to string) error {
	fmt.Println("***** Initiating from EsewaPayment *****", amount, from, to)
	return nil
}

// IMPLEMENTING DI USING CONSTRUCTOR
type OrderService struct {
	payment Payment
}

func NewOrderService(payment Payment) *OrderService {
	return &OrderService{
		payment: payment,
	}
}

func (o *OrderService) checkout(amount int, from string, to string) {
	o.payment.Payment(amount, from, to)
}

func main() {
	// Decide ONCE at startup which provider to use
	khalti := &KhaltiPayment{
		secret_key: "xxx",
		public_key: "yyy",
		live_key:   "zzz",
	}

	// Inject into the service
	service := NewOrderService(khalti)
	service.checkout(5000, "Ram", "Shiva")

	// Want Esewa instead? Swap
	esewa := &EsewaPayment{merchant_id: "CCC", secret_key: "VBB"}
	orderServiceEsewa := NewOrderService(esewa)
	orderServiceEsewa.checkout(5000, "Ram", "Shiva")

	fmt.Println("Exited!!!")

}
