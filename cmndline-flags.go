/* cmndline-flags.go
- specifies cmnd line program options.
*/

package main

import (
	"flag"
	"fmt"
)

func main() {

	// flag declarations are available for strings, integers
	// & booleans. this ex is for a string (returns a pointer)

	wordPtr := flag.String("word", "foo", "a string")

	// numb & fork flags

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// declare an option used elsewhere.
	// note the need to pass a pointer.

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// execute cmnd line parsing

	flag.Parse()

	// dump parsed options + any trailing positional args
	// note need to de-reference the pointers to get values

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

/* build first for experimentation:

$go build cmndline-flags.go

first try: using all defined flags
$./cmndline-flags -word=opt -numb=7 -fork -svar=flag

omitted flags ==> default values used
$./cmndline-flags -word=opt

trailing positional args
$./cmndline-flags -word=opt a1 a2 a3

use -h or --help for auto-generated help text
$./cmndline-flags -h

unrecognized flags ==> error message + help message
$./cmndline-flags -wat
*/
