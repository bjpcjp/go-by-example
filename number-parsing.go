/* number-parsing.go
- for parsing numbers from strings.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	// ParseFloat: "64" = #bits of precision to parse.

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// ParseInt: "0" = infer base from string.
	// "64" = result fits in 64 bits.

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt recognizes hex formats.

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// ParseUnit is available.

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi: convenience function for base 10 int parsing.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parse functions return errors on bad inputs.

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
