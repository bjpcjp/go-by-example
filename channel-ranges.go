/* channel-ranges.go
- using "for" and "range" to iterate over values
  received from a channel
*/

package main

import "fmt"

func main() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    for elem := range queue {
        fmt.Println(elem)
    }
}