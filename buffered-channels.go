/* channel-buffering.go
- channels are unbuffered by default (ie, they only accept
  sends) if there is a corresponding receive (<- chan) 
  ready to receive the value.

- buffered channels accept a limited number of values
  without a corresponding receiver.
*/

package main

import "fmt"

func main() {

    messages := make(chan string, 2)

    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
