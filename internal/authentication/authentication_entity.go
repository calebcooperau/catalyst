package authentication

import (
	"time"

	"github.com/google/uuid"
)

type AuthUser struct {
	ID           uuid.UUID
	Email        string
	FirstName    string
	LastName     string
	MobileNumber *string
}

var AnonymousUser = &AuthUser{}

func (usr *AuthUser) IsAnonymous() bool {
	return usr == AnonymousUser
}

func Create(email string, firstName string, lastName string) (*AuthUser, error) {
	user := AuthUser{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
	return &user, nil
}

type AuthProvider struct {
	ID             uuid.UUID
	UserID         uuid.UUID
	Provider       string
	ProviderUserID string
	CreatedAt      time.Time
}

func (authProvider AuthProvider) Create(userID uuid.UUID, provider string, providerUserID string) (*AuthProvider, error) {
	newProvider := &AuthProvider{
		UserID:         userID,
		Provider:       provider,
		ProviderUserID: providerUserID,
	}
	return newProvider, nil
}
