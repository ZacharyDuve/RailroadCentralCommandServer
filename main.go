package railroadcentralcommandserver

import (
	"log/slog"
	"os"
)

func main() {
	// Going to log to std out for now but best long term would be a centralized logger

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	logger.Info("Starting Railroad Central Command Server")
}
