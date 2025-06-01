package db

import "github.com/LoaltyProgramm/quotes-service/internal/models/quotes"

func InitStorage() map[int64]quotes.Quote {
	quotes := make(map[int64]quotes.Quote, 0)
	return quotes
}