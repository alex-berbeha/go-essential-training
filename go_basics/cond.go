// Example of "if" statement

package main

import "fmt"

func main() {
	x := 38
	if x > 35 {
		fmt.Println("x is big")
	}

	if x > 100 {
		fmt.Println("x is a very big")
	} else {
		fmt.Println("x is not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	if x < 5 || x > 15 {
		fmt.Println("x is out of range")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}
}

