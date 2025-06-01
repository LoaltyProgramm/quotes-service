package main

import (
	"log"

	"github.com/LoaltyProgramm/quotes-service/internal/config"
	"github.com/LoaltyProgramm/quotes-service/internal/db"
	"github.com/LoaltyProgramm/quotes-service/internal/handlers"
	"github.com/LoaltyProgramm/quotes-service/internal/repository"
	"github.com/LoaltyProgramm/quotes-service/internal/server"
	quoteservice "github.com/LoaltyProgramm/quotes-service/internal/services/quote_service"
)

func main() {
	cfg := config.NewConfig()
	quoteStorage := db.InitStorage()
	repo := repository.NewRepository(quoteStorage)
	quoteService := quoteservice.NewQuoteService(repo)
	handlers := handlers.NewHandlers(quoteService)

	handlers.InitHandlers()

	if err := server.RunServer(cfg); err != nil {
		log.Fatal(err)
	}
}
