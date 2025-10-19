package application

import (
	"log"
	"net/http"
	"os"
	"time"

	"catalyst.api/config"
	"catalyst.api/internal/database"
	"catalyst.api/internal/domain"
	"catalyst.api/internal/middleware"
	"catalyst.api/migrations"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

type Application struct {
	Logger       *log.Logger
	Gin          *gin.Engine
	Config       *config.Config
	Server       *http.Server
	Database     *pgxpool.Pool
	Repositories *domain.Repositories
	Middlewares  *middleware.Middlewares
}

func NewApplication(cfg *config.Config) (*Application, error) {
	pool, err := database.OpenDBPool()
	if err != nil {
		return nil, err
	}

	sqlDB := stdlib.OpenDBFromPool(pool)
	defer sqlDB.Close()

	err = database.MigrateFS(sqlDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	if cfg.HttpConfig.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Allow frontend URL (Angular app)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
	}))

	repositories := domain.RegisterRepositories(pool)
	middlewares := middleware.RegisterMiddlewares(repositories)
	server := &http.Server{
		Addr:         cfg.HttpConfig.Port,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      engine,
	}
	app := &Application{
		Logger:       logger,
		Gin:          engine,
		Config:       cfg,
		Server:       server,
		Database:     pool,
		Repositories: repositories,
		Middlewares:  middlewares,
	}

	return app, nil
}

func (app *Application) Start() {
	err := app.Server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
