/* line filters
- line filter: reads input from stdin, processes it,
  prints derived result to stdout.
- grep, sed = common line filters.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// example: writes capitalized version of all input text.

func main() {

	// wrapping unbuffered os.Stdin with buffered scanner
	// is convenient Scan method that advances scanner to
	// next token (which is the next line in the sccanner)

	scanner := bufio.NewScanner(os.Stdin)

	// Text returns current token from the input

	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text())

		// write the uppercased line

		fmt.Println(ucl)
	}

	// check for errors. EOL is expected and not
	// reported as an error.

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// try with:
// $echo 'hello' > /tmp/lines
// $echo 'filter' >> /tmp/lines
// $cat /tmp/lines | go run line-filters.go
