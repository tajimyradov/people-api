package main

import (
	"log"

	"github.com/joho/godotenv"
	"people-api/config"
	_ "people-api/docs"
	"people-api/internal/handler"
	"people-api/internal/repository"
	"people-api/internal/server"
	"people-api/internal/service"
	"people-api/pkg/logger"
	"people-api/pkg/storage"
)

// @title People API
// @version 1.0
// @description This is a sample People API.
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading sample.env file")
	}

	cfg := config.Load()

	db, err := storage.NewPostgres(cfg)
	if err != nil {
		log.Fatal("connect DB: ", err)
	}
	defer db.Close()

	enricher := service.NewEnricher(cfg)
	repo := repository.NewPersonRepository(db)
	personService := service.NewPersonService(repo, enricher)
	personHandler := handler.NewPersonHandler(personService)

	r := server.NewRouter(personHandler)

	logger.Log.Info("Starting server on port ", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
