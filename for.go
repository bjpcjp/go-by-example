/*
"for" is Go's only looping construct. There are 3 types:
1) basic: single condition
2) classic: initial/condition/after
3) forever: without a condition, for will loop repeatedly
   until a break or a return from an enclosing function.
"continue" goes to next iteration of the loop.
*/

package main

import "fmt"

func main() {

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}