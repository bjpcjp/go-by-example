/* functions.go */

package main

import "fmt"

// basics: accepts two ints, returns int sum

func plus(a int, b int) int {

	// go requires explicit returns
	// no auto-return of last expression

	return a + b
}

// multi params of the same type:
// you can omit the type name up to the last param

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
