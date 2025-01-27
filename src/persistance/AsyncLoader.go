package persistance

import (
	"context"

	"github.com/google/uuid"
)

type AsyncLoader[T Persistable] interface {
	// LoadAll returns a read only channel of Persistable items as well as errors
	// Success is when the channel of items closes and error is never returned
	// If an error is sent down the err channel then all reads from items is invalid from then on
	// Ctx is passed in to allow for early termination of the loading
	LoadAllAsync(ctx context.Context) (items <-chan T, err <-chan error)

	LoadByPrimaryKeyAsync(pk uuid.UUID, ctx context.Context) (items <-chan T, err <-chan error)
}
