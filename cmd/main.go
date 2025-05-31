package main

import (
	"log"

	"github.com/LoaltyProgramm/quotes-service/internal/config"
	//"github.com/LoaltyProgramm/quotes-service/internal/db"
	"github.com/LoaltyProgramm/quotes-service/internal/server"
)

func main() {
	//QuotesStorage := db.InitStorage()
	cfg := config.NewConfig()

	if err := server.RunServer(cfg); err != nil {
		log.Fatal(err)
	}
}
