package cache

import (
	"cmp"
	"errors"
	"fmt"
	"hash/maphash"
	"sync"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/storage"
)

type shardedCache[I cmp.Ordered, T storage.Storable[I]] struct {
	hashSeed maphash.Seed
	shards   []cacheShard[I, T]
}

type cacheShard[I cmp.Ordered, T storage.Storable[I]] struct {
	shardLock sync.Mutex
	items     []T
}

func NewShardedCache[I cmp.Ordered, T storage.Storable[I]](numShards int) (storage.Storage[I, T], error) {
	if numShards < 1 {
		return nil, errors.New("invalid number of shards (less than 1), unable to create ShardedCache")
	}

	sC := &shardedCache[I, T]{
		hashSeed: maphash.MakeSeed(),
		shards:   make([]cacheShard[I, T], numShards),
	}

	return sC, nil
}

func (sc *shardedCache[I, T]) Save(t T) error {
	index := sc.calcShardIndex(t.ID())

	return sc.shards[index].save(T)
}

func (sc *shardedCache[I, T]) Load(I) (T, error) {
	var t T = *new(T)
	return t, errors.ErrUnsupported
}

func (sc *shardedCache[I, T]) Delete(I) error {
	return errors.ErrUnsupported
}

func (sc *shardedCache[I, T]) calcShardIndex(id I) int {
	hash := maphash.Hash{}
	hash.SetSeed(sc.hashSeed)

	hash.WriteString(fmt.Sprint(id))

	return int(hash.Sum64()) % len(sc.shards)
}

func (cs *cacheShard[I, T]) save(t T) error {
	cs.shardLock.Lock()

	cs.shardLock.Unlock()

	return nil
}
