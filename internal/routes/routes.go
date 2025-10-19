package routes

import (
	"log"

	"catalyst.api/cmd/docs"
	"catalyst.api/internal/authentication"
	"catalyst.api/internal/domain"
	"catalyst.api/internal/domain/user"
	"catalyst.api/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(engine *gin.Engine, db *pgxpool.Pool, repos *domain.Repositories, middlewares *middleware.Middlewares, logger *log.Logger) {
	router := engine
	docs.SwaggerInfo.BasePath = "/"
	router.Use(middlewares.AuthenticationMiddleware.Authenticate())
	{
		authentication.RegisterRoutes(router, repos.AuthenticationRepository, logger)
		user.RegisterRoutes(router, db, repos.UserRepository, middlewares.AuthenticationMiddleware, logger)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
