gist-playground
===============

gist-playground is a commandline tool that allows you to turn a gist into a
[go playground][playground] sandboxed code snippet.

Getting it
----------

Set up your `$GOPATH` and install the latest version of [go][go], then
run the following:

````
go get github.com/ecnahc515/gist-playground/cmd/gistpl
````

You should now have the `gistpl` command which will allow you to provide a
gist-id (`http://gist.github.com/<user>/<id>`) and will return a
http://play.golang.org url.

Caveats
-------

gist-playground currently only supports retrieving a single file/package from a
gist. It currently looks for the first file in a gist with the contents
`package main` in it and uses that for creating a playground link. This is a
restriction of the [go playground][playground] more than gist-playground itself.


[playground]: http://play.golang.org/ "go playground"
[go]: http://golang.org/dl/