/* sorting.go 
- go has a "sort" package for built-in & user-defined datatypes.
- sort methods are type-specific. this example is for strings.
*/

package main

import "fmt"
import "sort"

func main() {

    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}