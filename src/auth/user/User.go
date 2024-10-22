package user

import "github.com/google/uuid"

type User struct {
	id    uuid.UUID
	name  string
	roles []UserRole
	store ObjectStore[*User]
}

func (this *User) ID() uuid.UUID {
	return this.id
}

func (this *User) Name() string {
	return this.name
}

func (this *User) ChangeName(nName string) error {
	//Want to check that the name is available first
	// Then save the user with the new name
	this.store.Update(this)
	return nil
}

type UserRole string
