/* string-formats.go */

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}

	// %v refers to "print verbs."
	fmt.Printf("%v\n", p)

	// include the struct's field names.
	fmt.Printf("%+v\n", p)

	// go syntax representation of the value.
	fmt.Printf("%#v\n", p)

	// return the value's datatype.
	fmt.Printf("%T\n", p)

	// booleans. straightforward.
	fmt.Printf("%t\n", true)

	// integers; std base-10.
	fmt.Printf("%d\n", 123)

	// integers: binary.
	fmt.Printf("%b\n", 14)

	// character equivalent of the integer.
	fmt.Printf("%c\n", 33)

	// hex encoding.
	fmt.Printf("%x\n", 456)

	// floats - basic decimal formatting.
	fmt.Printf("%f\n", 78.9)

	// scientific notation.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// basic strings.
	fmt.Printf("%s\n", "\"string\"")

	// double-quoting strings, as in Go source.
	fmt.Printf("%q\n", "\"string\"")

	// strings as base16 - 2 chars per input byte.
	fmt.Printf("%x\n", "hex this")

	// pointer
	fmt.Printf("%p\n", &p)

	// numbers with spec'd width & precision
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// width of printed floats
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// left-justified with "-"
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// strings with spec'd width, right-justified.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// strings with spec'd width, left-justified.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// format & return string without printing
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// printing to io.Writers other than os.Stdout
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
