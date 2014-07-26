package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ecnahc515/gist-playground/gist"
	"github.com/google/go-github/github"
)

var token string

func init() {
	token = os.Getenv("GISTPLAYGROUND_TOKEN")
}

func main() {
	flag.Parse()
	args := flag.Args()

	httpClient := gist.NewCachingHttpClient(token, nil, nil)
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
		gst, _, err := client.Gists.Get(id)
		if err != nil {
			fmt.Println("Error retrieving gist:", err.Error())
			os.Exit(1)
		}
		var content string
		content, err = gist.FindMain(gst)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(content)
	}
}
