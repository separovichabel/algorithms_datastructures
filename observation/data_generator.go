package observation

import (
	"math/rand"
	"time"
)

var negative = -100
var positive = 100

func RandSlice(len int) []int {
	randons := make([]int, len)

	for i := 0; i < len; i++ {
		randons[i] = randPositiveNegative()
	}

	return randons
}

func randPositiveNegative() int {
	rand.Seed(time.Now().UnixNano())
	rng := positive - negative
	return rand.Intn(rng) + negative
}
