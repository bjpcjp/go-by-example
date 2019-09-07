/* http-servers.go */

package main

import (
	"fmt"
	"net/http"
)

// fundamental net/http concept: handlers (an object implementing
// the http.Handler interface).

func hello(w http.ResponseWriter, req *http.Request) {

	// handler functions take a ResponseWriter & http.Request
	// as arguments.

	fmt.Fprintf(w, "hello\n")
}

// read all HTTP request headers; echo them into response body

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// register handlers on server routes using http.HandleFunc
// (a convenience function). it sets a default router & takes
// a function as an argument.

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// ListenAndServe with port# and handler.
	// nil = use default router we just defined.

	http.ListenAndServe(":8090", nil)
}

/* running server in the background:
   $go run http-servers.go &

   access /hello:
   $curl localhost:8090/hello
*/
