/* if-else.go

- you can have an if statement without an else.
- a statement can precede conditionals. any variables
  declared in this statement are available in all branches.
- you don't need parens around conditions, but braces are required.
- there is no ternary if in go. you need a full if statement
  even for basic conditions.
*/

package main

import "fmt"

func main() {

    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}