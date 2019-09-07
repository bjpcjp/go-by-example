/* errors.go
- go communicates errors with explicit return values
- this is different than Java & Ruby's exceptions.
- errors are the last return value & have an error datatype.
- errors.new constructs an error value with a given msg.
- nil in the error position = no error.
- you can use custom types as errors via Error().

- more info:
http://blog.golang.org/2011/07/error-handling-and-go.html

*/

package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {

		return -1, errors.New("can't work with 42")

	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
