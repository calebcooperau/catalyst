package main

import (
	"catalyst.api/config"
	"catalyst.api/internal/application"
	"catalyst.api/internal/authentication"
	"catalyst.api/internal/routes"
)

// @title catalyst.api
// @version 1.0
// @description Generate Project Management Tasks from Diagrams
// @host localhost:42069
// @BasePath /
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	authentication.NewAuthentication(cfg)
	app, err := application.NewApplication(cfg)
	if err != nil {
		panic(err)
	}
	defer app.Database.Close()
	routes.SetupRoutes(app.Gin, app.Database, app.Repositories, app.Middlewares, app.Logger)
	app.Start()
}
