package server

import (
	"log"
	"net/http"

	"github.com/ecnahc515/goplayutils/gist"
	"github.com/ecnahc515/goplayutils/playground"
	"github.com/google/go-github/github"
)

type Server struct {
	Client *github.Client
}

func (server *Server) indexHandler(rw http.ResponseWriter, req *http.Request) {
	gistid := req.FormValue("gistid")
	if gistid != "" {
		log.Println("Received request for gistid:", gistid)
		content, err := gist.GetGist(server.Client, gistid)
		if err != nil {
			http.Error(rw, err.Error(), 500)
			return
		}
		url, err := playground.GetPlayUrl(&content)
		if err != nil {
			http.Error(rw, err.Error(), 500)
			return
		}
		rw.Write([]byte(url))
	}
}

func (server *Server) registerHandlers() {
	http.HandleFunc("/gist", server.indexHandler)
}

func (server *Server) Start(addr string) {
	server.registerHandlers()
	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
