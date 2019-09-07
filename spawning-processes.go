/* spawning-processes.go */

package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	// simple cmnd - takes no arguments, just prints to stdout.
	// Command: create object to represent the external process.

	dateCmd := exec.Command("date")

	// Output: handles running cmnd, waiting for finish,
	// and collecting output.

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// more complicated example:
	// pipe data to external process on its stdin,
	// collect results from its stdout.

	grepCmd := exec.Command("grep", "hello")

	// explicitly grab input & output pipes

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	// start the process
	grepCmd.Start()

	// write to it, close the grepIn, read result output,
	// wait for process to exit.

	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// error checks not done above. you can use
	// "if err != nil" to get them.

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// when spawning commands, need to provide
	// explicit command & argument array
	// to spawn a full cmnd with a string,
	// use bash's -c option.

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

/*
$go run spawning-processes.go
>date
>grep hello
>ls -a -l -h
*/
