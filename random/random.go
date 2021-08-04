package random

import (
	"math/rand"
)

type Random struct {
	Source int64
}

func (r *Random) HundredNumber() int {
	rand.Seed(r.Source)
	n := rand.Intn(100)
	return n
}
