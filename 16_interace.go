package main

import "fmt"

// Interface: Payment provider must implement Pay()
type Payment interface {
	Pay(amount float64)
}

// UPI payments
type UPI struct{}

func (UPI) Pay(amount float64) {
	fmt.Println("Paid using UPI:", amount)
}

// Credit Card
type Card struct{}

func (Card) Pay(amount float64) {
	fmt.Println("Paid using Credit Card:", amount)
}

// Function accepts ANY payment method
func Checkout(p Payment, amount float64) {
	p.Pay(amount)
	// Pay is present in both the struct
}

func main() {
	upi := UPI{}
	card := Card{}

	// Same function works with different types
	Checkout(upi, 500.0)   // Paid via UPI
	Checkout(card, 1000.0) // Paid via Card
}
