/* ranges.go */

package main

import "fmt"

func main() {

	// using range to sum numbers in a slice.

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range on arrays & slices provides the index
	// & value for each entry.

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range on maps - iterates over key:value pairs

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range iterating over just the keys of a map

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
