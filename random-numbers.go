/* random-numbers.go */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// rand.Intn returns integer (0..100, in this case)

	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// rand.Float returns a Float64 (0.0..1.0)
	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// default generator is deterministic - produces
	// same sequence each time by default. use a changing
	// seed to generate varying sequences.

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// if you seed with the same number, it produces
	// the same random sequence.

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
