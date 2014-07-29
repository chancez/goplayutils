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

`goplay` can take input in a few ways. You can provide it a `gist-id` which is
the last segment after your username in a github gist. If your providing a
specific revision, everything after your username should be used.

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

#### As a Web Service:

The other option is to run it as a web service which accepts http requests and
returns a [playground][playground] url. To this you can run a command such as:

````
goplay -d --addr "0.0.0.0:8080"
````

The `-d` flag means to run as a web service, and the `--addr` flag is optional
and lets you specify a non default host/port.

Once it's running you can access the webpage at the root of localhost or the
fqdn of the host its running on. It also accepts requests with a content type of
`application/json` and a `gistid` query parameter at the `/gist` endpoint. It
will then return the url in the JSON form `{"url": "play.golang.org/p/####"}`.

If you have goplay running on `localhost:8080` (the default) you could run the
following:

````
curl -H "Content-Type: application/json" "http://localhost:8080/gist?gistid=952190cba18de244b472"
{"url":"http://play.golang.org/p/bvnGZ_Uf7g"}
````

Caveats
-------

goplayutils currently only supports retrieving a single file/package from a
gist. It currently looks for the first file in a gist with the contents
`package main` in it and uses that for creating a playground link. This is a
restriction of the [go playground][playground] more than goplayutils itself.


[playground]: http://play.golang.org/ "go playground"
[go]: http://golang.org/dl/