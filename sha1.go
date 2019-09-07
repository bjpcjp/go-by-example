/* sha1.go
- SHA1 hashes often used to compute short IDs for binary
  or text blobs.
- can also generate MD5 hashes using crypto/md5 and
  md5.New().
- also: hash strengh:
  http://en.wikipedia.org/wiki/Cryptographic_hash_function
*/

package main

// uses crypto/sha1 package.

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	// generates new hash

	h := sha1.New()

	// Write expects bytes. use []byte(s) to coerce
	// strings to bytes if needed.

	h.Write([]byte(s))

	// final hash result as a byte slice

	bs := h.Sum(nil)

	// SHA1 vals often printed in hex (ex: Git commits).
	// use %x format to convert results to a hex string.

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
