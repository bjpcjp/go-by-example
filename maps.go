/* maps.go
- maps are go's built-in associative data type (aka hashes or dicts)
- create an empty map, use "make(map[keytype]valtype)
- set key:val pairs with name[key]=value.
- get a key value with name[key].
- len returns the number of key:val pairs
- delete removes key:val pairs
- optional 2nd return value indicates whether the given
  key was present in the map. this can be ignored via
  a blank identifier "_".

*/

package main

import "fmt"

func main() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}