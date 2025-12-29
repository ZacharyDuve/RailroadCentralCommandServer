package cache

import "github.com/ZacharyDuve/RailroadCentralCommandServer/src/storage"

type shardedCache[T storage.Storable] struct {
	shards []cacheShard[T]
}

type cacheShard[T storage.Storable] struct {
}

func NewShardedCache[T storage.Storable](numShards int) {

}
