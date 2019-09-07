/* channels.go
- channels are pipes between concurrent goroutines.
- you can send values into channels from one goroutine
  and get those values from another.

- send a value to a channel with channel <-
- receive a value from a channel with <- channel

- by default, sends & receives block until both the
  sender and receiver are ready.
*/

package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
