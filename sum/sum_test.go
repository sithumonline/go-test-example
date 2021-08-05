package sum

import (
	"testing"

	"github.com/sithumonline/go-test-example/number"
)

type Test struct {
	meth TestX
}

type TestX interface {
	GetNumber() int
}

func (s *Test) GetNumber() int {
	n := 184
	return n
}

type Tos struct {
	meth TosX
}

type TosX interface {
	SetRandomNumber(int64)
	GetRandomNumber() int
}

func (s *Tos) SetRandomNumber(f int64) {}

func (s *Tos) GetRandomNumber() int {
	n := 2184
	return n
}

func TestGetSum(t *testing.T) {
	y := NewSum(&Test{})

	m := y.GetSum(100)
	if m != 284 {
		t.Error("Expected: 284, got:", m)
	}

	t.Logf("Got MockHundredNumber : %d", m)

	n := number.NewNumber(&Tos{})
	n.LoadNumber()

	z := NewSum(n)

	u := z.GetSum(200)
	if u != 2384 {
		t.Error("Expected: 2284, got:", u)
	}

	t.Logf("Got MockThousandHundredNumber : %d", u)
}
