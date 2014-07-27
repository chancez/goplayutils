Goplayutils
===============

goplayutils is a command line tool that allows you to turn gists and other go
code into a [go playground][playground] sandboxed code snippet.

Getting it
----------

Set up your `$GOPATH` and install the latest version of [go][go], then
run the following:

````
go get github.com/ecnahc515/goplayutils/cmd/goplay
````

You should now have the `goplay` binary in your `$GOPATH`.

Usage
-----

`goplay` can take input in a few ways. You can provide it a `gist-id` which is the last segment after your username in a github gist. If your providing a specific revision,
everything after your username should be used.

#### Examples:

The gist: https://gist.github.com/ecnahc515/952190cba18de244b472

A specific revision of that gist: https://gist.github.com/ecnahc515/952190cba18de244b472/865f92da4ab99d404d10dc5b705461332b7687a1

````
# no revision
goplay 952190cba18de244b472
# with the revision:
goplay 952190cba18de244b472/865f92da4ab99d404d10dc5b705461332b7687a1
````

You can also use goplay on non-gists by passing in the file contents via stdin:

````
cat somefile.go | goplay
````

Caveats
-------

goplayutils currently only supports retrieving a single file/package from a
gist. It currently looks for the first file in a gist with the contents
`package main` in it and uses that for creating a playground link. This is a
restriction of the [go playground][playground] more than goplayutils itself.


[playground]: http://play.golang.org/ "go playground"
[go]: http://golang.org/dl/