package storage

import (
	"os"
	"testing"
)

type mockData struct {
	id uint
}

func (m *mockData) ID() uint {
	return m.id
}

func (m *mockData) TypeName() string {
	return "mockData"
}

type mockFileManager struct {
	FileManager
}

func (mFM *mockFileManager) MkdirAll(p string, m os.FileMode) error {
	return nil
}

func TestJSONStorageImplementsStorage(t *testing.T) {
	s, _ := NewJSONStorage[uint, *mockData]("bob", &mockFileManager{})

	switch s.(type) {
	case Storage[uint, *mockData]:
		return
	}
	t.Fail()
}
