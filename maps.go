/* maps.go

maps = built-in associative data type
(aka hashes or dicts) */

package main

import "fmt"

func main() {

	// create an empty map with make

	m := make(map[string]int)

	// set key:value pairs

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// get value from a key

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// len

	fmt.Println("len:", len(m))

	// delete key:value pair

	delete(m, "k2")
	fmt.Println("map:", m)

	// optional 2nd return value - indicates whether
	// key is present in a map
	// can ignore value itself with "_"

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
