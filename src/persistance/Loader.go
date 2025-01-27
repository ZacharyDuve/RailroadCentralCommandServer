package persistance

import "github.com/google/uuid"

type Loader[T Persistable] interface {
	LoadAll() ([]T, error)
	LoadByPrimaryKey(pk uuid.UUID) (T, error)
}
