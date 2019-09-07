/* panic.go
- usually used to quickly fail on unexpected errors
  (esp if not prepared to handle gracefully).
- below: an example panicking if we get an expected
  error when creating a new file.
*/

package main

import "os"

func main() {

    panic("a problem")

    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}