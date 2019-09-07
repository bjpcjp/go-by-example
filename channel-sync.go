/* channel-sync.go
- we can use channels to sync execution across goroutines.
- if waiting for multiple goroutines to finish, consider
  using WaitGroups instead.
*/

package main

import "fmt"
import "time"

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    <-done
}