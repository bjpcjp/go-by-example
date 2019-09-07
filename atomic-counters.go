/* atomic-counters.go
- the primary mechanism for managing state in Go is
  comms over channels.
- we can also use atomic counters via the sync/atomic
  package.
*/

package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {

    var ops uint64

    var wg sync.WaitGroup

    for i := 0; i < 50; i++ {
        wg.Add(1)

        go func() {
            for c := 0; c < 1000; c++ {

                atomic.AddUint64(&ops, 1)
            }
            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Println("ops:", ops)
}