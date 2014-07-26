package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ecnahc515/gist-playground/gist"
	"github.com/ecnahc515/gist-playground/playground"
	"github.com/google/go-github/github"
	"github.com/gregjones/httpcache/diskcache"
)

const cacheDir = "gist_playground_cache"

var token string

func init() {
	token = os.Getenv("GISTPLAYGROUND_TOKEN")
	flag.Usage = Usage
}

func NewDiskCache() *diskcache.Cache {
	tmpDir := os.TempDir()
	path := filepath.Join(tmpDir, cacheDir)
	return diskcache.New(path)
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Error: Specify a gist ID to use.\n")
		fmt.Println("example: run 'gp 952190cba18de244b472'")
		os.Exit(1)
	}

	switch args[0] {
	case "serve":
		// set up http server
	default:
		id := args[0]
		cache := NewDiskCache()
		httpClient := gist.NewCachingHttpClient(token, cache, nil)
		client := github.NewClient(httpClient)
		content, err := gist.GetGist(client, id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		playUrl, err := playground.GetPlayUrl(&content)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(playUrl)
	}
}
