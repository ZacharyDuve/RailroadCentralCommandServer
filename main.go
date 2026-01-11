package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/apirest"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	logger.Info("Starting Railroad Central Command Server")

	http.ListenAndServe(":8080", apirest.ServeRESTAPIs())

}
