package main

import (
	"fmt"
	"os"

	"github.com/ecnahc515/gist-playground/gist"
	"github.com/ecnahc515/gist-playground/playground"
	"github.com/google/go-github/github"
)

func printPlayUrlId(id string) {
	cache := NewDiskCache()
	httpClient := gist.NewCachingHttpClient(token, cache, nil)
	client := github.NewClient(httpClient)
	content, err := gist.GetGist(client, id)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	printPlayUrl(&content)
}

func printPlayUrl(content *string) {
	playUrl, err := playground.GetPlayUrl(content)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(playUrl)

}
