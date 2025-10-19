package authentication

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, authenticationRepo AuthenticationRepository, logger *log.Logger) {
	// Set up handlers
	signInHandler := NewSignInHandler(authenticationRepo, logger)
	logoutHandler := NewLogoutHandler(logger)
	providerHandler := NewProviderHandler(logger)

	// Set up routes
	authRoutes := router.Group("/auth")
	authRoutes.GET("/:provider/callback", signInHandler.SignInCallback)
	authRoutes.GET("/logout/:provider", logoutHandler.Logout)
	authRoutes.GET("/:provider", providerHandler.GetProvider)
}
