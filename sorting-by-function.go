/* sorting-by-functions.go
- how to sort by something other than natural order.
  (for example, strings by length, not alphabetically).
*/

package main

import (
	"fmt"
	"sort"
)

// alias for built-in []string datatype
type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

/* by using this pattern (creating a custom type,
   creating 3 interface methods for the type,
   then calling sort.Sort on a collection of that type),
   we can sort Go slices with arbitrary functions.
*/

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
