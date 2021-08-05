package random

import (
	"math/rand"
)

type Random struct {
	Source int64
}

func NewRandom() *Random {
	return &Random{}
}

func (r *Random) SetRandomNumber(f int64) {
	r.Source = f
}

func (r *Random) GetRandomNumber() int {
	rand.Seed(r.Source)
	n := rand.Intn(100)
	return n
}
