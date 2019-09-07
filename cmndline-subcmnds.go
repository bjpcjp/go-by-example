/* cmndline-subcmnds.go */

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// declare subcmnds using NewFlagSet

	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// different subcmnds = separate flags

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// subcmnd is expected to be 1st arg to the program

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	// for each subcmnd,
	// parse its own flags & have access to trailing args

	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

/*
$go build cmndline-subcmnds.go

# first invocation
$./cmndline-subcmnds foo -enable -name=joe a1 a2

# try bar
$./cmndline-subcmnds bar -level 8 a1

# bar doesn't accept foo's flags
$./cmndline-subcmnds bar -enable a1

*/
