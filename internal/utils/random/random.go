package random

import (
	"math/rand"
	"time"
)

func RandomInt(n int) int {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	randomNumb := rnd.Intn(n)

	return randomNumb
}


