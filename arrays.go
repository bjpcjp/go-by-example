/* arrays.go */

package main

import "fmt"

func main() {

	// create array to hold exactly 5 integers.
	// element types & length are declared.
	// default arrays are zero-valued.

	var a [5]int
	fmt.Println("emp:", a)

	// normal get & set array terminology

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// len

	fmt.Println("len:", len(a))

	// declare & initialize array in one statement

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// arrays are 1D; you can build multi-D structures.

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
