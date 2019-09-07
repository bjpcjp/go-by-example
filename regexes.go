/* regexes.go */

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// tests whether a pattern matches a string.

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// sometimes you'll need to compile a construct
	// instead of directly using a string pattern.

	r, _ := regexp.Compile("p([a-z]+)ch")

	// basic match.

	fmt.Println(r.MatchString("peach"))

	// also finds 1st match, but returns the start & end
	// indexes instead of the matching text.

	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))

	// return sub-match info

	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// non-negative integer as 2nd arg will
	// limit the number of matches.

	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// byte arguments.

	fmt.Println(r.Match([]byte("peach")))

	// use MustCompile when creating constants with regexes.
	// plain Compile won't work for constants - it returns
	// 2 values.

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// replace subsets of strings with other values.

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Func = allows you to transform matched text
	// with a given function.

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
