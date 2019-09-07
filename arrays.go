/* arrays.go
- array element types & length of part of the definition.
- default fill value is zero.
- array[index]=value; array[index] returns value.
- len: returns array length.
- arrays are 1D; you can compose types to build multi-D
  data structures.
- slices are much more common than arrays in Go.
*/

package main

import "fmt"

func main() {

    var a [5]int
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}