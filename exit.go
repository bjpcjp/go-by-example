/* exit.go */

package main

import (
	"fmt"
	"os"
)

func main() {

	// defers will NOT be run when using Exit,
	// so this Println will never be called.

	defer fmt.Println("!")

	// exit with status=3

	os.Exit(3)
}

/*
$go run exit.go --> exit status 3
$go build exit.go
$./exit --> binary build lets you see status in the terminal
$echo $? --> 3
*/
