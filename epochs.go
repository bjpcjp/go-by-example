/* epochs.go
- common need: getting the number of seconds, msecs
  or nanosecs since the Unix epoch:
  http://en.wikipedia.org/wiki/Unix_time
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	// get elapsed time
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// no UnixMillis - have to build it yourself.

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// converting integer secs or nanosecs since the epoch
	// into a corresponding time.

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
