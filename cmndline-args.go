/* cmndline-args.go */

package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Args can access raw command-line args.
	// 1st value = program path
	// os.Args[1:] = program arguments.

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// use normal indexing for individual args

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

/* for experimentation, it's best to build a binary first:

$go build cmndline-args.go
$./cmndline-args a b c d
[a b c d]
c
*/
