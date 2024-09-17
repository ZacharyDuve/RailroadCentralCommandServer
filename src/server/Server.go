package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/server/api"
)

type ServerConfig interface {
	PortNumber() uint16
}

// Would like to have some sort of config to get passed in
func ListenAndServeServer(conf ServerConfig) {
	//Prep the handlers
	mux := http.DefaultServeMux

	mux.Handle("/api", &api.ApiHandler{})
	mux.Handle("/", http.FileServer(http.Dir("./web")))

	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", conf.PortNumber()), mux))
}
