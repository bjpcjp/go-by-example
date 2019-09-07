/* variables.go
- variables are explicitly declared & used by the compiler
  to check the datatypes in function calls.
- var declares 1 or more variables.
- uninitialized vars are zero-valued.
*/

package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
