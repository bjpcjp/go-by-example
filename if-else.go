/* if-else.go */

package main

import "fmt"

func main() {

	// basic example

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// if without an else clause

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// statements can precede conditions,
	// variables declared here are available
	// in all branches

	// you don't need parentheses around conditions,
	// but braces are required.

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
