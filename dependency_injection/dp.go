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

// Implementing interface :-> Method level DI
// This is the Dependency Inversion Principle from SOLID. MakingPayment depends on the Payment interface (abstraction), not on Khalti or Esewa (concrete implementations).
func MakingPayment(payment Payment, amount int, from string, to string) {
	payment.Payment(amount, from, to)
}

func main() {
	khalti := &KhaltiPayment{
		secret_key: "xxxxxxxxxxxxxxxxxxxxxxxxxx",
		public_key: "yyyyyyyyyyyyyyyyyyyyyyyyyy",
		live_key:   "zzzzzzzzzzzzzzzzzzzzzzzzzzz",
	}

	esewa := &EsewaPayment{
		merchant_id: "CCCCCC",
		secret_key:  "VBBBB",
	}

	MakingPayment(khalti, 5000, "Ram", "Shiva")
	MakingPayment(esewa, 5000, "Ram", "Shiva")

}
