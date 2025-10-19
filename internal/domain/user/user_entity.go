package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Email        string
	FirstName    string
	LastName     string
	MobileNumber string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func Create(email string, firstName string, lastName string) *User {
	user := &User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}

	return user
}

func (usr *User) Update(email string, firstName string, lastName string, mobile string) (*User, error) {
	err := usr.CanUpdate()
	if err != nil {
		return nil, err
	}

	usr.Email = email
	usr.FirstName = firstName
	usr.LastName = lastName
	usr.MobileNumber = mobile
	// need to additional update for password
	return usr, nil
}

func (usr *User) CanUpdate() error {
	// add in validation
	return nil
}

func (usr *User) CanDelete() error {
	// add in can delete validation
	return nil
}
