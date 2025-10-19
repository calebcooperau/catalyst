package domain

import (
	"catalyst.api/internal/authentication"
	"catalyst.api/internal/domain/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	UserRepository           user.UserRepository
	AuthenticationRepository authentication.AuthenticationRepository
}

func RegisterRepositories(db *pgxpool.Pool) *Repositories {
	userRepository := user.NewUserSqlRepository(db)
	authenticationRepository := authentication.NewAuthenticationSqlRepository(db)
	return &Repositories{
		UserRepository:           userRepository,
		AuthenticationRepository: authenticationRepository,
	}
}
