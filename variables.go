/* variables.go
- variables are explicitly declared & used by the compiler
  to check the datatypes in function calls.
- uninitialized vars are zero-valued.
*/

package main

import "fmt"

func main() {

	// var declares one or more variables.

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// go infers variable datatypes.

	var d = true
	fmt.Println(d)

	// uninitialized variables are zero-valued.

	var e int
	fmt.Println(e)

	// ":=" is shorthand for declaring & initializing
	// variables.

	f := "apple"
	fmt.Println(f)
}
