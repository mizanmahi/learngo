package main

import "fmt"

// func makeCoffee(kind string) (string, int) { // multiple return values
// 	price := 25
// 	coffee := fmt.Sprintf("%s Coffee!", kind)
// 	return coffee, price
// }

func makeCoffee(kind string) (coffee string, price int) { // named return value
	price = 25
	coffee = fmt.Sprintf("%s Coffee!", kind)
	return
}

func main() {
	myCoffee, myBill := makeCoffee("black")
	myCoffee2, myBill2 := makeCoffee("cold")

	fmt.Printf("I am having a %d$ %s \n", myBill, myCoffee)
	fmt.Printf("I am having a %d$ %s", myBill2, myCoffee2)

}
