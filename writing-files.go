/* writing-files.go */

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// start: dump a string to a files

	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// for more granular writes - open a file for writing

	f, err := os.Create("/tmp/dat2")
	check(err)

	// customary to defer a Close right after opening a file.

	defer f.Close()

	// Writing byte slices

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// WriteString

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// Sync flushes writes to stable storage

	f.Sync()

	// buffered writers

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// Flush ensures all buffered ops are applied

	w.Flush()

}
