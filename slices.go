/* slices.go
more details at:
http://blog.golang.org/go-slices-usage-and-internals
*/

package main

import "fmt"

func main() {

	// use make to create empty non-zero-length slices

	s := make([]string, 3)
	fmt.Println("emp:", s)

	// getters, setters & len, just like arrays

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// append

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice operator: [low:high]

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// declare & initialize in a single statement

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// multi-D slice creation

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
