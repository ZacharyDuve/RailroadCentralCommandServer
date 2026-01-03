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
	shardLock sync.RWMutex
	items     map[I]T
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

	return sc.shards[index].save(t)
}

func (sc *shardedCache[I, T]) Load(id I) (T, error) {
	index := sc.calcShardIndex(id)

	return sc.shards[index].load(id)
}

func (sc *shardedCache[I, T]) Delete(id I) error {
	index := sc.calcShardIndex(id)

	return sc.shards[index].delete(id)
}

func (sc *shardedCache[I, T]) calcShardIndex(id I) int {
	hash := maphash.Hash{}
	hash.SetSeed(sc.hashSeed)

	hash.WriteString(fmt.Sprint(id))

	return int(hash.Sum64()) % len(sc.shards)
}

func (cs *cacheShard[I, T]) save(t T) error {
	cs.shardLock.Lock()

	cs.items[t.ID()] = t

	cs.shardLock.Unlock()

	return nil
}

func (cs *cacheShard[I, T]) load(id I) (T, error) {

	cs.shardLock.RLock()

	t, ok := cs.items[id]

	cs.shardLock.RUnlock()

	if !ok {
		return t, fmt.Errorf(storage.ErrFmtMsgUnableToLoadDoesNotExist, id)
	}

	return t, nil
}

func (cs *cacheShard[I, T]) delete(id I) error {
	cs.shardLock.Lock()

	_, ok := cs.items[id]

	if !ok {
		return fmt.Errorf(storage.ErrFmtMsgUnableToDeleteDoesNotExist, id)
	} else {
		delete(cs.items, id)
	}

	cs.shardLock.Unlock()

	return nil
}
