/* temp-files.go */

// often need to create temp data. temp files & directories
// can help, since they don't pollute the os over time.

package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

	// TempFile: easiest way to create a temp file
	// "" as 1st argument = create file in default location.

    f, err := ioutil.TempFile("", "sample")
    check(err)

	// display temp file's name. directory likely to be
	// "/tmp" on Unix systems.

    fmt.Println("Temp file name:", f.Name())

    defer os.Remove(f.Name())

	// write file data

    _, err = f.Write([]byte{1, 2, 3, 4})
    check(err)

	// if writing many temp files, consider creating
	// a temp directory

    dname, err := ioutil.TempDir("", "sampledir")
    fmt.Println("Temp dir name:", dname)

    defer os.RemoveAll(dname)

	// synthesize temp file names by using the 
	// temp directory as a prefix
	
    fname := filepath.Join(dname, "file1")
    err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
    check(err)
}