package main

import (
	"log/slog"
	"os"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/storage/persistent"
)

type PersonID uint

type PersonJSON struct {
	IDField   PersonID `json:"id"`
	FirstName string   `json:"first-name"`
	LastName  string   `json:"last-name"`
}

func (pj PersonJSON) ID() PersonID {
	return pj.IDField
}

func main() {
	// Going to log to std out for now but best long term would be a centralized logger

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	logger.Info("Starting Railroad Central Command Server")

	personStorage, err := persistent.NewJSONStorage[PersonID, *PersonJSON]("./data", persistent.OSFileManager{})

	if err != nil {
		logger.Error("error occurred trying to create personStorage", "error", err)
		os.Exit(1)
	}

	p := PersonJSON{IDField: 124, FirstName: "Jerry", LastName: "Nguyen"}

	if err := personStorage.Save(&p); err != nil {
		logger.Error("error occurred trying to save person", "person", p, "error", err)
	} else {
		logger.Info("successfully saved person", "person", p)
	}

	// if p2, err := personStorage.Load(123); err != nil {
	// 	logger.Error("error occurred trying to save person", "error", err)
	// } else {
	// 	logger.Info("successfully loaded person", "person", p2)

	// }

	if err := personStorage.Delete(123); err != nil {
		logger.Error("error occurred trying to save person", "error", err)
	} else {
		logger.Info("successfully deleted person")

	}

}
