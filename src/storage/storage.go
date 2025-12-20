package storage

import "github.com/ZacharyDuve/RailroadCentralCommandServer/src/id"

// Storage is the interface that describes something that can store and retrieve an object. Object must implement id.IDable
type Storage[K id.IDable] interface {
	// Save attempts to save K as the storage defines.
	// An error could occur and be returned.
	Save(*K) error

	// Load attempts to load K back from storage.
	// On success K is returned with nil error
	// Failure then error is returned
	Load(id.ID) (*K, error)
}
