package track

import "github.com/ZacharyDuve/RailroadCentralCommandServer/src/storage"

type TrackID uint32

type Track struct {
	id           TrackID
	name         string
	lengthMeters float32
	speedKPH     float32
}

func (t *Track) ID() TrackID {
	return t.id
}

type TrackManager struct {
	trackStore storage.Storage[TrackID, *Track]
}
