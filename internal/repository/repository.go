package repository

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/LoaltyProgramm/quotes-service/internal/models/quotes"
	"github.com/LoaltyProgramm/quotes-service/internal/utils/random"
)

type Repository interface {
	AddQuote(quote quotes.Quote) error
	GetQuotes() (map[int64]quotes.Quote, error)
	GetQuoteRandom() (map[int64]quotes.Quote, error)
	GetQuotesByAuthor(author string) (map[int64]quotes.Quote, error)
	DeleteQuote(id string) error
}

type repository struct {
	Storage    map[int64]quotes.Quote
	counterIdx int64
	mu         sync.Mutex
}

func NewRepository(storage map[int64]quotes.Quote) Repository {
	return &repository{
		Storage:    storage,
		counterIdx: 0,
	}
}

func (r *repository) AddQuote(quote quotes.Quote) error {
	id := atomic.AddInt64(&r.counterIdx, 1)

	for _, v := range r.Storage {
		if v == quote {
			return errors.New("such a record already exists")
		}
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.Storage[id] = quote

	return nil
}

func (r *repository) GetQuotes() (map[int64]quotes.Quote, error) {
	quotes := make([]quotes.Quote, 0, 10)
	for _, v := range r.Storage {
		quotes = append(quotes, v)
	}

	if len(quotes) == 0 {
		return nil, errors.New("quotes is not found")
	}

	return r.Storage, nil
}

func (r *repository) GetQuoteRandom() (map[int64]quotes.Quote, error) {
	keys := make([]int64, 0, 10)
	for k := range r.Storage {
		keys = append(keys, k)
	}

	if len(keys) == 0 {
		return nil, errors.New("quotes is not found")
	}

	lenKey := len(keys)

	randomIdx := random.RandomInt(lenKey)

	key := keys[randomIdx]

	randomQuote := r.Storage[key]

	return map[int64]quotes.Quote{ // Тут правка была
		key: randomQuote,
	}, nil
}

func (r *repository) GetQuotesByAuthor(author string) (map[int64]quotes.Quote, error) {
	authorQuotes := make(map[int64]quotes.Quote)
	for k, v := range r.Storage {
		if v.Author == author {
			authorQuotes[k] = quotes.Quote{
				Author: v.Author,
				Quote: v.Quote,
			}
		}
	}

	if len(authorQuotes) == 0 {
		return nil, fmt.Errorf("author is not found")
	}

	return authorQuotes, nil
}

func (r *repository) DeleteQuote(id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	idInt64 := int64(idInt)

	if _, ok := r.Storage[idInt64]; !ok {
		return errors.New("no record was found for this id")
	}

	delete(r.Storage, idInt64)

	return nil
}
