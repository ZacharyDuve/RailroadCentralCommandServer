package user

import "github.com/google/uuid"

type Identifiable interface {
	ID() uuid.UUID
}

type userStore interface {
	Create(u *User) error
	Update(u *User) error
	DeleteById(u uuid.UUID) error
	GetById(u uuid.UUID) (*User, error)
}

type ObjectStore[T Identifiable] interface {
	Create(u T) error
	Update(u T) error
	DeleteById(u uuid.UUID) error
	GetById(u uuid.UUID) (T, error)
}
