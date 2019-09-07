/* base64.go */

package main

// this syntax imports encoding/base64 with the b64 name
// of the default base64.

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// here's the string we'll encode/decode.

	data := "abc123!?$*&()'-=@~"

	// Go supports std & URL-compatible base64.
	// here: encode using std encoder.

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// decoding can return an error.
	// check if you aren't sure if input is well-formed.

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// encode/decode using URL-compatible base64

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
