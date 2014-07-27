package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	cliUsage = "<gist id>"
	webUsage = "[options]"
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Cli Usage: %s %s\n", os.Args[0], cliUsage)
	fmt.Fprintf(os.Stderr, "Webapp Usage: %s %s\n", os.Args[0], webUsage)
	fmt.Fprint(os.Stderr, "Options:\n")
	flag.PrintDefaults()
}
