/* multi-return-values.go */

package main

import "fmt"

// (int, int) indicates two integers will be returned.

func vals() (int, int) {
	return 3, 7
}

func main() {

	// multiple assignment
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// use blank identifier if you only need a
	// subset of the returned values

	_, c := vals()
	fmt.Println(c)
}
