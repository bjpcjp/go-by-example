/* environment-vars.go */

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// use Setenv to set a key:value pair.
	// use Getenv to get a value for a key.

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// use Environ to list all key:value pairs
	// returns a slice of strings - use string.Split
	// to get individual values

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}

/* $go run environment-vars.go
   $BAR=2 go run environment-vars.go
*/
