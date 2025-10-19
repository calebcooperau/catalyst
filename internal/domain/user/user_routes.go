package user

import (
	"log"

	"catalyst.api/internal/authentication"
	"catalyst.api/internal/domain/user/data"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(router *gin.Engine, db *pgxpool.Pool, repo UserRepository, authMiddleware authentication.AuthenticationMiddleware, logger *log.Logger) {
	queries := data.New(db)
	// Set up handlers
	detailHandler := NewUserDetailHandler(queries, logger)
	updateHandler := NewUserUpdateHandler(repo, logger)
	deleteHandler := NewUserDeleteHandler(repo, logger)

	// Set up routes
	userRoutes := router.Group("/user")
	userRoutes.Use(authMiddleware.RequireAuthUser())
	{
		userRoutes.GET("/:id", detailHandler.GetUserByID)
		userRoutes.PUT("/:id", updateHandler.UpdateUser)
		userRoutes.DELETE("/:id", deleteHandler.DeleteUser)
	}
}
