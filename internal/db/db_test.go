package db

import (
	"reflect"
	"testing"

	"github.com/LoaltyProgramm/quotes-service/internal/models/quotes"
)

func TestInitStorage(t *testing.T) {
	quotesMemory := InitStorage()

	expectedType := reflect.TypeOf(make(map[int64]quotes.Quote))
	actualType := reflect.TypeOf(quotesMemory)

	if quotesMemory == nil {
		t.Errorf("InitStorage() returned nil map, want initialized map")
	}

	if actualType != expectedType {
		t.Errorf("InitStorage() returned wrong type: got %v, want %v", actualType, expectedType)
	}
}