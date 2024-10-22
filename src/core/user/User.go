package user

import "github.com/ZacharyDuve/RailroadCentralCommandServer/src/persistance"

const (
	TYPENAME_USER string = "USER"
)

type User struct {
	userName string
	persistance.StorableTimeStamp
}

func (this *User) PrimaryKey() persistance.PrimaryKey {
	return persistance.PrimaryKey(this.userName)
}

func (this *User) TypeName() string {
	return TYPENAME_USER
}

func (this *User) UserName() string {
	return this.userName
}

type Credentials struct {
}
