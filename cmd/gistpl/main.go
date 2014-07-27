package main

import (
	"flag"
	"fmt"

	"os"
	"path/filepath"

	"github.com/ecnahc515/gist-playground/server"
	"github.com/gregjones/httpcache/diskcache"
)

const cacheDir = "gist_playground_cache"

var (
	token  string
	daemon = flag.Bool("d", false, "Enable to daemonize and run as a web app.")
	addr   = flag.String("addr", "localhost:8080", "Address and port to listen on.")
)

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

	if *daemon {
		if len(args) > 0 {
			fmt.Println("-d flag requires no arguments.")
			os.Exit(1)
		}
		server.Start(*addr)
	}

	if len(args) < 1 {
		fmt.Printf("Error: Specify a gist ID to use.\n\n")
		fmt.Printf("example: run '%s 952190cba18de244b472'\n", os.Args[0])
		os.Exit(1)
	}

	handleCli(args)
}
