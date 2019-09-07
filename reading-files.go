/* reading-files.go */

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// file reads require error checking.
// this helper will help below.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// basic read op:
	// slurping entire file contents into memory.

	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// for more control; start with opening file
	// to obtain os.File value.

	f, err := os.Open("/tmp/dat")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// seek to a known location in the file

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// ReadAtLeast

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// Seek(0,0) = effectively a "rewind" function

	_, err = f.Seek(0, 0)
	check(err)

	// bufio.NewReader = buffered reader

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// close when done.

	f.Close()
}
