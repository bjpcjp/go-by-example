/* multi-return-values.go
- go supports multiple return values. often used for
  returning both result & error values from a function.
- blank identifiers accepted.
*/

package main

import "fmt"

func vals() (int, int) {
    return 3, 7
}

func main() {

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals()
    fmt.Println(c)
}