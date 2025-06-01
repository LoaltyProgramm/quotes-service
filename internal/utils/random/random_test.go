package random

import (
	"testing"
)

func TestRandomInt(t *testing.T) {
	id := RandomInt(10)
	
	if id > 10 {
		t.Errorf("the number is greater than the specified specified interval")
	}

	if id < 0 {
		t.Errorf("the number cannot be less than zero")
	}
}