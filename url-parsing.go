/* url-parsing.go */

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// test URL: includes
	// scheme, auth info, host, port, path, query info.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// parse the url. ensure no errors.

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// accessing the scheme

	fmt.Println(u.Scheme)

	// user contains all authentication info
	// Username, Password = individual values.

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// Host contains hostname & port if present.

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// extract path & fragment after the "#".

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// use RawQuery to get params in a k=v string

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
