/* collections.go
- we often need to do operations on data collections,
  ex selecting all items that satisfy a condition.
- below: example collection functions on string slices.
*/

package main

import (
	"fmt"
	"strings"
)

// index:
// return 1st index of t, or -1 if not found.

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// include:
// returns true if t is in the slice.

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// any:
// returns true if any string in the slice satisfies t.

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// all:
// returns true if all strings in the slice satisfy t.

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// filter:
// returns new slice with all strings in slice
// that satisfy t.

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// map:
// returns new slice with results of applying function f
// to each string in the slice.

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// try'em all out.

func main() {

	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))

}
