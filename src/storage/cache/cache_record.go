package cache

import (
	"cmp"
	"time"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/storage"
)

type cacheRecord[I cmp.Ordered, T storage.Storable[I]] struct {
	lastAccessTime time.Time
	hits           uint64
	t              T
}
