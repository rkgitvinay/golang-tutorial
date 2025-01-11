//In a payment processing system, different payment methods such as CreditCard, Upi, and BankTransfer
// share similar behaviours like Pay and Refund.
// Using interfaces, we can model these payment methods and allow polymorphic behaviour.

package main

import "fmt"

// Define the interface
type PaymentProcessor interface {
	Pay(amount float64) string
	Refund(amount float64) string
}

// Implement the Interface for Different Payment Methods

// Credit Card
type CreditCard struct {
	CardNumber string
}

func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card ending with %s", amount, cc.CardNumber[len(cc.CardNumber)-4:])
}

func (cc CreditCard) Refund(amount float64) string {
	return fmt.Sprintf("Refunded %.2f using Credit Card ending with %s", amount, cc.CardNumber[len(cc.CardNumber)-4:])
}

// Upi
type Upi struct {
	Mobile string
}

func (up Upi) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using UPI number: %s", amount, up.Mobile)
}

func (up Upi) Refund(amount float64) string {
	return fmt.Sprintf("Refunded %.2f to UPI number: %s", amount, up.Mobile)
}

// Bank Transfer
type BankTransfer struct {
	BankName string
	Account  string
}

func (bt BankTransfer) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f via Bank Transfer to %s, Account: %s", amount, bt.BankName, bt.Account)
}

func (bt BankTransfer) Refund(amount float64) string {
	return fmt.Sprintf("Refunded %.2f via Bank Transfer to %s, Account: %s", amount, bt.BankName, bt.Account)
}

// Use the Interface in a Function
func ProcessPayment(p PaymentProcessor, amount float64) {
	fmt.Println(p.Pay(amount))
}

func RefundPayment(p PaymentProcessor, amount float64) {
	fmt.Println(p.Refund(amount))
}

func main() {
	// Create payment method instances
	creditCard := CreditCard{CardNumber: "1234567890"}
	upi := Upi{Mobile: "8285846281"}
	bankTransfer := BankTransfer{BankName: "Icici", Account: "123456789"}

	// Proces payments
	fmt.Println("Processing Payments:")
	ProcessPayment(creditCard, 100.00)
	ProcessPayment(upi, 200.73)
	ProcessPayment(bankTransfer, 33.22)

	// Process Refund
	fmt.Println("\nProcessing Refunds:")
	RefundPayment(creditCard, 50.25)
	RefundPayment(upi, 100.00)
	RefundPayment(bankTransfer, 300.00)
}
