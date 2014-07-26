package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `<gist id>
`

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s %s\n", os.Args[0], usage)
	flag.PrintDefaults()
}
