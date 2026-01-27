package main

import "fmt"

<<<<<<< HEAD
func main() {
	sugar := 2

	makeCoffee := func() {
		coffee := "Cappuccino"
		sugar := 3 // modify
		fmt.Printf("Making %s with %d spoon of sugar \n", coffee, sugar)
		fmt.Println("Value of inner sugar", sugar)
=======
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
>>>>>>> 48324f5 (update main.go to implement named return values in makeCoffee function and demonstrate usage in main)

	}

	makeCoffee()

	fmt.Println("Value of outer sugar", sugar)
}
