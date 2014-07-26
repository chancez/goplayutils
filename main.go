package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/go-github/github"
	"github.com/gregjones/httpcache"
	"github.com/sourcegraph/apiproxy"
	"github.com/sourcegraph/apiproxy/service/github"
)

func NewCachingHttpClient(trans *httpcache.Transport,
	validator *apiproxy.Validator) *http.Client {

	if trans == nil {
		trans = httpcache.NewMemoryCacheTransport()
	}
	var check apiproxy.Validator
	if validator == nil {
		age := &githubproxy.MaxAge{
			User:         time.Hour * 24,
			Repository:   time.Hour * 24,
			Repositories: time.Hour * 24,
			Activity:     time.Hour * 12,
		}
		check = age.Validator()
	} else {
		check = *validator
	}
	transport := &apiproxy.RevalidationTransport{
		Transport: trans,
		Check:     check,
	}
	return &http.Client{Transport: transport}
}

func main() {
	flag.Parse()
	args := flag.Args()

	httpClient := NewCachingHttpClient(nil, nil)
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
