/* for.go */

package main

import "fmt"

func main() {

	// basic construct - single condition

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic initial/condition/after construct

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without a condition loops repeatedly.
	// use break to exit, or
	// return from enclosing function

	for {
		fmt.Println("loop")
		break
	}

	// use continue to go to next iteration

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
