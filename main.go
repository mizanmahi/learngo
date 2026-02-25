package main // executable package

import (
	"fmt"
	"learngo/payment"
	"learngo/test"
)

func init() {
	fmt.Println("Initializing db...")
}

func print() {
	fmt.Println("Im in test function")
}

func main() {

	// bkash := Bkash{apiKey: "sdfsfd"}
	// paymentService := NewPaymentService(bkash)
	// paymentService.checkout()

	// nagad := NewNagad("safsfsf")

	mockPm := test.MockPaymentMethod{}

	paymentService1 := payment.NewPaymentService(mockPm)
	paymentService1.Checkout()

	// color.Red("Prints text in red.")
	// color.BgRGB(255, 128, 0).Println("background orange")

	fmt.Println("Im in main function")

	print()

}

// module = bunch of packages
