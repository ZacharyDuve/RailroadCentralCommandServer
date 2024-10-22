package persistance

import (
	"time"
)

type PrimaryKey string

type Storable interface {
	//Name of the type that we are trying to store. Allows for some genericish of data storage
	TypeName() string
	//Key that is supposed to be unique that uniquely identifies this one object
	PrimaryKey() PrimaryKey
	//Needs to have some sort of timestamp as part of the object to help control
	TimeStampable
}

type Storer[T Storable] interface {
	Add(T) error
	Update(T) error
	Delete(PrimaryKey) error
	FindByPK(PrimaryKey) (T, error)
}

type TimeStampable interface {
	CreateTime() time.Time
	UpdateTime() time.Time
	RefreshUpdateTime()
}

type StorableTimeStamp struct {
	createTime time.Time
	updateTime time.Time
}

func (this *StorableTimeStamp) Init() {
	now := time.Now()
	this.createTime = now
	this.updateTime = now
}

func (this *StorableTimeStamp) CreateTime() time.Time {
	return this.createTime
}

func (this *StorableTimeStamp) UpdateTime() time.Time {
	return this.updateTime
}

func (this *StorableTimeStamp) RefreshUpdateTime() {
	this.updateTime = time.Now()
}
