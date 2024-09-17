package main

import (
	"log"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/config"
	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/server"
)

func main() {
	log.Println("Starting Railroad-Central-Command-Server, have fun training")
	server.ListenAndServeServer(&config.Config{})
}
