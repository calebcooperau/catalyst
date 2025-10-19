package middleware

import (
	"catalyst.api/internal/authentication"
	"catalyst.api/internal/domain"
)

type Middlewares struct {
	AuthenticationMiddleware authentication.AuthenticationMiddleware
}

func RegisterMiddlewares(repositories *domain.Repositories) *Middlewares {
	authenticationMiddleware := authentication.AuthenticationMiddleware{AuthenticationRepository: repositories.AuthenticationRepository}

	middlewares := &Middlewares{
		AuthenticationMiddleware: authenticationMiddleware,
	}

	return middlewares
}
