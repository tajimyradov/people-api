package main

import (
	"flag"
	"go.elastic.co/apm/module/apmzap"
	"go.uber.org/zap"
	"log"

	"github.com/joho/godotenv"
	"people-api/config"
	_ "people-api/docs"
	"people-api/internal/handler"
	"people-api/internal/repository"
	"people-api/internal/server"
	"people-api/internal/service"
	ll "people-api/pkg/logger"
	"people-api/pkg/storage"
)

const (
	Name = "people-api"
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

	var logLevel string
	flag.StringVar(&logLevel, "ll", "info", "logging level")

	db, err := storage.NewPostgres(cfg)
	if err != nil {
		log.Fatal("connect DB: ", err)
	}
	defer db.Close()

	logger, err := ll.NewWithSampler(
		Name,
		logLevel,
		1,
		500,
		zap.WrapCore((&apmzap.Core{}).WrapCore),
	)
	if err != nil {
		log.Fatal("error while init logger", err.Error())
	}

	logger.Info(
		"flags",
		zap.String("log_level", logLevel),
	)

	enricher := service.NewEnricher(cfg)
	repo := repository.NewPersonRepository(db)
	personService := service.NewPersonService(repo, enricher, logger)
	personHandler := handler.NewPersonHandler(personService)

	r := server.NewRouter(personHandler)

	logger.Info("Starting server on port ", zap.String("port", cfg.ServerPort))
	r.Run(":" + cfg.ServerPort)
}
