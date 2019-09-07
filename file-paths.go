/* file-paths.go */

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// use Join to construct portable paths.

	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// use Join instead of concatenating /s or \s manually.

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Dir & Base split paths to directories & files.
	// Split is another alternative.

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// Check if a path is absolute.

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// some file names have extensions following a dot.
	// split the extension with Ext.

	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// TrimSuffix gets filename with extension removed.

	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel finds relative path between base & target.

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
