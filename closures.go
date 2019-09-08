/* closures.go
- go supports anonymous functions, which form closures.
*/

package main

import "fmt"

// intSeq's internal anonymous function
// forms a closure around the variable i.

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
