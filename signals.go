/* signals.go

example: we might want a server to gracefully shut down
when it receives a SIGTERM, or a command-line tool to
stop processing input if it receives a SIGINT.

more info: http://en.wikipedia.org/wiki/Unix_signal */

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// go signals work by sending Signal values on a channel.
	// create a channel to get these notifications, another
	// to notify us when done.

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Notify registers the channel to receive notifications.

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// goroutine: executes a blocking receive for signals.
	// when it gets one, print it & notify program that
	// it can finish.

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// program waits here until signal is received,
	// then exit

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

/*
$go run signals.go -- control-C to send SIGINT signal,
causing program to print interrupt, then exit.
*/
