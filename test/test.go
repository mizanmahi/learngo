package test

import "fmt"

func init() {
	fmt.Println("Im from test package init function")
}

type MockPaymentMethod struct {
}

func (mockPM MockPaymentMethod) Pay(amount float64) {
	// test logic
	fmt.Println("testing payment method")
}
