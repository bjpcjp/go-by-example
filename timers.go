/* timers.go
- "timer" and "ticker" allow us to execute code at some
  point in the future, or on intervals.
- Timers represent a single future event.
- you can also use time.Sleep; timers allow you to 
  cancel the time before it expires.
*/

package main

import "time"
import "fmt"

func main() {

    timer1 := time.NewTimer(2 * time.Second)

    <-timer1.C
    fmt.Println("Timer 1 expired")

    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}