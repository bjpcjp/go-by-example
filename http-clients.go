/* http-clients.go */

package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// issue HTTP GET request to a server
	// Get = shortcut around creating a Client object
	// and calling its Get method.

	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// response status

	fmt.Println("Response status:", resp.Status)

	// first five lines of response body

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
