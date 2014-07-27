package server

import (
	"log"
	"net/http"
)

func Start(addr string) {
	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
