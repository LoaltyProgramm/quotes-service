package db

import "github.com/LoaltyProgramm/quotes-service/internal/models/quotes"

func InitStorage() map[string]quotes.Quotes {
	quotes := make(map[string]quotes.Quotes, 0)
	return quotes
}