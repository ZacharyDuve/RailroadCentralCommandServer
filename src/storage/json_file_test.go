package storage

import (
	"io"
	"testing"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/compare"
)

type mockData struct {
	d int
}

func (m *mockData) ID() int {
	return m.d
}

func (m *mockData) Compare(o *mockData) compare.CompareResult {
	if m.d < o.d {
		return compare.LessThan
	} else if m.d > o.d {
		return compare.GreaterThan
	} else {
		return compare.Equal
	}
}

func (m *mockData) Equal(o *mockData) bool {
	return m.d == o.d
}

func mockOpenFile(r io.ReadWriteCloser, err error) (io.ReadWriteCloser, error) {
	return r, err
}

func TestJSONStorageImplementsStorage(t *testing.T) {
	var _ Storage[*mockData] = NewJSONStorage[int]("hey", "int")
}
