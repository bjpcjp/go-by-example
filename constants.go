/* constants.go
- go supports character, string, boolean & numeric constants.
- consts can appear anywhere a variable can.
- const expressions do math with arbitrary precision.
- numeric constants do not have a type until given one.
- numbers can be given a type by using it in a context
  that requires one, eg math.Sin expects a float64.
*/

package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
