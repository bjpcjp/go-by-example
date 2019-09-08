/* constants.go */

package main

import (
	"fmt"
	"math"
)

// const = constant declaration.

const s string = "constant"

func main() {
	fmt.Println(s)

	// constants can be declared anywhere a variable can.

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	// numeric constants have no type until given one,
	// such as by explicit conversion.

	fmt.Println(int64(d))

	// numbers can be typed by using it in a context that
	// requires one.

	fmt.Println(math.Sin(n))
}
