package repository

import (
	"testing"

	"github.com/LoaltyProgramm/quotes-service/internal/models/quotes"
)

func TestAddQuote(t *testing.T) {
	t.Run("Unused data", func(t *testing.T) {
		storage := make(map[int64]quotes.Quote)

		repo := &repository{
			Storage:    storage,
			counterIdx: 0,
		}

		quote := quotes.Quote{
			Author: "Author1",
			Quote:  "Quote1",
		}

		err := repo.AddQuote(quote)
		if err != nil {
			t.Error(err.Error())
		}

		if len(storage) <= 0 {
			t.Errorf("expected 1 quote, got %d", len(storage))
		}

		for k, v := range repo.Storage {
			if v != quote {
				t.Errorf("expected %v, got %v", quote, v)
			}

			if k != 1 {
				t.Errorf("expected 1, got %d", k)
			}
		}
	})

	t.Run("used data", func(t *testing.T) {
		storage := map[int64]quotes.Quote{
			1: {Author: "Author1", Quote: "Quote1"},
		}

		repo := &repository{
			Storage:    storage,
			counterIdx: 0,
		}

		quote := quotes.Quote{
			Author: "Author1",
			Quote:  "Quote1",
		}

		expectedError := "such a record already exists"

		err := repo.AddQuote(quote)
		if err == nil {
			t.Errorf("expected error != nil, got %v", err)
		}

		if err.Error() != expectedError {
			t.Errorf("expected %v, got %v", expectedError, err.Error())
		}
	})

}

func TestGetQuotes(t *testing.T) {
	t.Run("The data is there", func(t *testing.T) {
		storage := map[int64]quotes.Quote{
			1: {Author: "Author1", Quote: "Quote1"},
			2: {Author: "Author2", Quote: "Quote2"},
			3: {Author: "Author3", Quote: "Quote3"},
		}

		repo := &repository{
			Storage: storage,
		}

		got, err := repo.GetQuotes()
		if err != nil {
			t.Error(err.Error())
		}

		expected := map[int64]quotes.Quote{
			1: {Author: "Author1", Quote: "Quote1"},
			2: {Author: "Author2", Quote: "Quote2"},
			3: {Author: "Author3", Quote: "Quote3"},
		}

		if len(got) != len(expected) {
			t.Fatalf("expected %d quotes, got %d", len(expected), len(got))
		}

		expectedSet := make(map[string]bool)
		for _, v := range expected {
			expectedSet[v.Quote] = true
		}

		for _, v := range got {
			if !expectedSet[v.Quote] {
				t.Errorf("unexpected quote: %v", v)
			}
		}
	})

	t.Run("No data available", func(t *testing.T) {
		storage := make(map[int64]quotes.Quote)

		repo := &repository{
			Storage: storage,
		}

		expectedError := "quotes is not found"

		quotes, err := repo.GetQuotes()
		if len(quotes) > 0 {
			t.Errorf("expected 0, got %v", len(quotes))
		}

		if err == nil {
			t.Errorf("expected error != nil, got %v", err)
		}

		if err.Error() != expectedError {
			t.Errorf("expected %v, got %v", expectedError, err.Error())
		}
	})
}

func TestGetQuoteRandom(t *testing.T) {
	t.Run("The data is there", func(t *testing.T) {
		storage := map[int64]quotes.Quote{
			1: {Author: "Author1", Quote: "Quote1"},
			2: {Author: "Author2", Quote: "Quote2"},
			3: {Author: "Author3", Quote: "Quote3"},
		}

		repo := &repository{
			Storage: storage,
		}

		randomQuote, err := repo.GetQuoteRandom()
		if err != nil {
			t.Error(err.Error())
		}

		valid := false
		for k, v := range storage {
			if randomQuote[k] == v {
				valid = true
				break
			}
		}

		if !valid {
			t.Errorf("expected true, got %v", valid)
		}
	})

	t.Run("No data available", func(t *testing.T) {
		storage := make(map[int64]quotes.Quote)

		repo := &repository{
			Storage: storage,
		}

		expectedError := "quotes is not found"

		randomQuote, err := repo.GetQuoteRandom()

		if len(randomQuote) > 0 {
			t.Errorf("expected 0, got %v", len(randomQuote))
		}

		if err == nil {
			t.Errorf("expected error != nil, got %v", err)
		}

		if err.Error() != expectedError {
			t.Errorf("expected %v, got %v", expectedError, err.Error())
		}
	})

}

func TestGetQuotesByAuthor(t *testing.T) {
	t.Run("Valid author", func(t *testing.T) {
		storage := map[int64]quotes.Quote{
			1: {Author: "Author1", Quote: "Quote1"},
			2: {Author: "Author2", Quote: "Quote2"},
			3: {Author: "Author3", Quote: "Quote3"},
			4: {Author: "Author2", Quote: "Quote4"},
			5: {Author: "Author4", Quote: "Quote5"},
			6: {Author: "Author2", Quote: "Quote6"},
		}

		repo := &repository{
			Storage: storage,
		}

		got, err := repo.GetQuotesByAuthor("Author2")
		if err != nil {
			t.Error(err.Error())
		}

		expected := map[int64]quotes.Quote{
			2: {Author: "Author2", Quote: "Quote2"},
			4: {Author: "Author2", Quote: "Quote4"},
			6: {Author: "Author2", Quote: "Quote6"},
		}

		if len(got) != len(expected) {
			t.Fatalf("expected len %d, got len %d", len(expected), len(got))
		}

		expectedSet := make(map[string]bool)
		for _, v := range expected {
			expectedSet[v.Author] = true
		}

		for _, v := range got {
			if !expectedSet[v.Author] {
				t.Errorf("unexpected quote %v", v)
			}
		}
	})

	t.Run("No valid Author", func(t *testing.T) {
		storage := map[int64]quotes.Quote{
			1: {Author: "Author1", Quote: "Quote1"},
			2: {Author: "Author2", Quote: "Quote2"},
			3: {Author: "Author3", Quote: "Quote3"},
			4: {Author: "Author2", Quote: "Quote4"},
			5: {Author: "Author4", Quote: "Quote5"},
			6: {Author: "Author2", Quote: "Quote6"},
		}

		repo := &repository{
			Storage: storage,
		}

		expectedError := "author is not found"

		_, err := repo.GetQuotesByAuthor("Author7")
		if err == nil {
			t.Errorf("expected error != nil, got %v", err)
		}

		if err.Error() != expectedError {
			t.Errorf("expected %v, got %v", expectedError, err.Error())
		}
	})

}

func TestDeleteQuote(t *testing.T) {
	t.Run("Valid id", func(t *testing.T) {
		storage := map[int64]quotes.Quote{
			1: {Author: "Author1", Quote: "Quote1"},
			2: {Author: "Author2", Quote: "Quote2"},
			3: {Author: "Author3", Quote: "Quote3"},
			4: {Author: "Author2", Quote: "Quote4"},
			5: {Author: "Author4", Quote: "Quote5"},
			6: {Author: "Author2", Quote: "Quote6"},
		}

		repo := &repository{
			Storage: storage,
		}

		err := repo.DeleteQuote("3")
		if err != nil {
			t.Fatalf("err not nil: %v", err)
		}

		for k := range storage {
			if k == 3 {
				t.Errorf("expected the index 3 should be deleted, got %d", k)
			}
		}
	})

	t.Run("Invalid id", func(t *testing.T) {
		storage := map[int64]quotes.Quote{
			1: {Author: "Author1", Quote: "Quote1"},
			2: {Author: "Author2", Quote: "Quote2"},
			3: {Author: "Author3", Quote: "Quote3"},
			4: {Author: "Author2", Quote: "Quote4"},
			5: {Author: "Author4", Quote: "Quote5"},
			6: {Author: "Author2", Quote: "Quote6"},
		}

		repo := &repository{
			Storage: storage,
		}

		expectedError := "no record was found for this id"

		err := repo.DeleteQuote("8")
		if err == nil {
			t.Errorf("expected error != nil, got %v", err)
		}

		if err.Error() != expectedError {
			t.Errorf("expected %v, got %v", expectedError, err.Error())
		}
	})

}
