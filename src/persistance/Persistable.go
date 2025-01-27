package persistance

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// TypeName is what you identify this persistable data as.
// Would recommend having the name line up to concret types instead of abstract. (structs vs interfaces)
// Example would be employee, vendor, car, etc.
type TypeName string

// PrimaryKey identifies each item of the same type from each other.
// type PrimaryKey []byte

// Persistable defines what a struct needs to implement to be persisted in some way.
type Persistable[K comparable] interface {
	TypeName() TypeName
	//PKey returns the primary key for the persisable item. This should be unique accross all of the items of the same type
	PKey() Key[K]
	CreateTime() time.Time
	UpdateTime() time.Time
	UpdateVersion() uint32
	MarkAsUpdated()
}

type PersistableMetaData struct {
	TName     TypeName
	UUIDField uuid.UUID
	CTime     time.Time
	UTime     time.Time
	UVersion  uint32
}

func (pMD *PersistableMetaData) Init() {
	if pMD.TName == TypeName("") {
		panic(errors.New("Unable to init PersistableMetaData due to TypeName field being unset"))
	}

	if pMD.UVersion != 0 {
		panic(errors.New("PersistableMetaData has already been initialized"))
	}

	now := time.Now()

	pMD.CTime = now
	pMD.UTime = now

	pMD.UUIDField = uuid.New()
}

func (pMD *PersistableMetaData) TypeName() TypeName {
	return pMD.TName
}

func (pMD *PersistableMetaData) UUID() uuid.UUID {
	return pMD.UUIDField
}

func (pMD *PersistableMetaData) CreateTime() time.Time {
	return pMD.CTime
}

func (pMD *PersistableMetaData) UpdateTime() time.Time {
	return pMD.UTime
}

func (pMD *PersistableMetaData) UpdateVersion() uint32 {
	return pMD.UVersion
}

func (pMD *PersistableMetaData) MarkAsUpdated() {
	pMD.UTime = time.Now()
	pMD.UVersion = pMD.UVersion + 1
}
