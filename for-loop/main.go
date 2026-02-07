package main

import "fmt"

func makeCoffee(coffeeNo int) {
	fmt.Println("making coffee....", coffeeNo)
}

func main() {

	for i := 1; i <= 5; i++ { // for initialization; condition; increment/decrement
		makeCoffee(i)
	}

	//

}

// i := 1, true, run the body, increment
// i = 2, true, run the body, increment
// i = 3, true, run the body, increment
// i = 4, true, run the body, increment
// i = 5, true, run the body, increment
// i = 6, false, =======
