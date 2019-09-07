/* directories.go */

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

	// create new subdirectory in current working directory.

	err := os.Mkdir("subdir", 0755)
	check(err)

	// when creating temp directories, its good practice to
	// defer their removal.
	// RemoveAll will delete entire directory tree. Use with
	// caution!

	defer os.RemoveAll("subdir")

	// helper function

	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// create hierarchy of directories with MkdirAll
	// similar to $mkdir -p

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// ReadDir lists directory contents. returns
	// slice of os.FileInfo objects.

	c, err := ioutil.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Chdir changes current working directory.

	err = os.Chdir("subdir/parent/child")
	check(err)

	// see contents of subdir/parent/child when
	// listing the current directory.

	c, err = ioutil.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// cd back to where we started

	err = os.Chdir("../../..")
	check(err)

	// visit a directory recursively including all subdirs.
	// Walk accepts a callback function to handle every
	// listed file or directory.

	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

// visit is called for every file/directory found by Walk

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
