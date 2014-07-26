package gist

import (
	"testing"

	"github.com/google/go-github/github"
)

func TestFindMainValid(t *testing.T) {
	pkg := `package main

import "fmt"

func main() {
    fmt.Println("hello world")
}`
	filename := "main.go"
	gst := &github.Gist{
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(filename): github.GistFile{
				Filename: &filename,
				Content:  &pkg,
			},
		},
	}

	content, err := FindMain(gst)
	if err != nil {
		t.Error(err)
	}
	if content != pkg {
		t.Errorf("FindMain(%v) = %v, want %v", gst, content, pkg)
	}

}

func TestFindMainFail(t *testing.T) {
	pkg := `package hello

import "fmt"

func main() {
    fmt.Println("hello world")
}`
	filename := "main.go"
	gst := &github.Gist{
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(filename): github.GistFile{
				Filename: &filename,
				Content:  &pkg,
			},
		},
	}

	content, err := FindMain(gst)
	if err != ErrNoPackageMain {
		t.Errorf("FindMain(%v) = (%v, %v), want %v", gst, content, err, "")
	}
	if content != "" {
		t.Errorf("FindMain(%v) = (%v, %v), want %v", gst, content, err, "")
	}

}
