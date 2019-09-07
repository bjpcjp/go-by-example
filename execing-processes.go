/* execing-processes.go

use this technique when you need an external process
accessible to a running Go process. sometimes you want
to replace a current Go process with another (perhaps
non-Go) one. */

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// example: execute "ls". Go needs an absolute path
	// to the binary. use LookPath to find it.
	// (probably /bin/ls.)

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec needs arguments to be a slice. Below:
	// some common example args.

	args := []string{"ls", "-a", "-l", "-h"}

	// Exec also needs a set of environment variables.

	env := os.Environ()

	// actual Exec call. if sucessful, this process will end
	// and be replaced by "/bin/ls -a -l -h". If an error,
	// we'll get a return value.

	// note: Go doesn't offer the classic Unix fork function.

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
