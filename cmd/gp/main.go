package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ecnahc515/gist-playground/gist"
	"github.com/google/go-github/github"
)

func main() {
	flag.Parse()
	args := flag.Args()

	httpClient := gist.NewCachingHttpClient(nil, nil)
	client := github.NewClient(httpClient)

	if len(args) < 1 {
		fmt.Println("Error, must provide at least one argument.")
		os.Exit(1)
	}

	switch args[0] {
	case "serve":
		// set up http server
	default:
		// passing in a url for a gist
		id := args[0]
		gist, _, err := client.Gists.Get(id)
		if err != nil {
			fmt.Println("Error retrieving gist:", err.Error())
			os.Exit(1)
		}
		fmt.Println(gist)
	}
}
