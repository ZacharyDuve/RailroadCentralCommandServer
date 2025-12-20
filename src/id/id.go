package id

// Key is a slice of bytes that contains a unique Key
type ID []byte

type IDable interface {
	Key()
}
