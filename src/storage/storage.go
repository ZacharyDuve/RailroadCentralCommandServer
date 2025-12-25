package storage

import "cmp"

// type SaveMode uint8

// const (
// 	// OverwriteExisting saves over the existing otherwise creates new.
// 	Overwrite SaveMode = iota
// 	// Create only saves if net new.
// 	Create
// 	// Update only saves if already exists.
// 	Update
// )

const (
	ErrFmtMsgUnableToSaveAlreadyExist   string = "unable to save %s as it already exists"
	ErrFmtMsgUnableToLoadDoesNotExist   string = "unable to load %s as it does not exist"
	ErrFmtMsgUnableToDeleteDoesNotExist string = "unable to delete %s as it does not exist"
)

type Storable[T cmp.Ordered] interface {
	ID() T
	TypeName() string
}

// Storage is the interface that describes something that can store and retrieve an object. Object must implement id.IDable
type Storage[I cmp.Ordered, T Storable[I]] interface {
	// Save attempts to save T as the storage defines. Currently either create or overwrite
	// An error could occur and be returned.
	Save(T) error

	// Load attempts to load T back from storage.
	// On success T is returned with nil error
	// Failure then error is returned
	Load(I) (T, error)

	// Delete attempts to delete T from storage so future loads will not find it.
	// Failure to delete returns error
	Delete(I) error
}
