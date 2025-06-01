package quoteservice

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/LoaltyProgramm/quotes-service/internal/models/quotes"
	"github.com/LoaltyProgramm/quotes-service/internal/repository"
)

type QuoteService interface {
	CreateQuote(quote quotes.Quote) error
	ListQuotes() ([]quotes.Quote, error)
	GetQuoteRandom() (*quotes.Quote, error)
	ListQuotesByAuthor(author string) ([]string, error)
	RemoveQuoteById(id string) error
}

type quoteService struct {
	repo repository.Repository
}

func NewQuoteService(repo repository.Repository) QuoteService {
	return &quoteService{
		repo: repo,
	}
}

func (s *quoteService) CreateQuote(quote quotes.Quote) error {
	if quote.Author == "" || quote.Quote == "" {
		return fmt.Errorf("author or quote is not empty")
	}

	err := s.repo.AddQuote(quote)
	if err != nil {
		return err
	}

	return nil
}

func (s *quoteService) ListQuotes() ([]quotes.Quote, error) {
	quotes, err := s.repo.GetQuotes()
	if err != nil {
		return nil, err
	}
	return quotes, nil
}

func (s *quoteService) GetQuoteRandom() (*quotes.Quote, error) {
	quote, err := s.repo.GetQuoteRandom()
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func (s *quoteService) ListQuotesByAuthor(author string) ([]string, error) {
	 _, err := strconv.Atoi(author)
	 if err == nil {
		return nil, errors.New("the data cannot be a number")
	 }

	 quotesByAuthor, err := s.repo.GetQuotesByAuthor(author)
	 if err != nil {
		return nil, err
	 }

	 return quotesByAuthor, nil
}

func (s *quoteService) RemoveQuoteById(id string) error {
	if _, err := strconv.Atoi(id); err != nil {
		return errors.New("the identifier cannot be a letter")
	}

	err := s.repo.DeleteQuote(id)
	if err != nil {
		return err
	}

	return nil
}