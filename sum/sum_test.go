package sum

import (
	"testing"

	"github.com/sithumonline/go-test-example/number"
)

type MockNumber struct {
	method MockNumberMethod
}

type MockNumberMethod interface {
	GetNumber() int
}

func (m *MockNumber) GetNumber() int { return 184 }

type MockRandom struct {
	method MockRandomMethod
}

type MockRandomMethod interface {
	SetRandomNumber(int64)
	GetRandomNumber() int
}

func (m *MockRandom) SetRandomNumber(f int64) {}

func (m *MockRandom) GetRandomNumber() int { return 2184 }

func TestGetSum(t *testing.T) {
	sum := NewSum(&MockNumber{}).GetSum(100)
	if sum != 284 {
		t.Error("Expected: 284, got:", sum)
	}

	t.Logf("Got MockHundredNumber : %d", sum)

	num := number.NewNumber(&MockRandom{})
	num.LoadNumber()

	sum = NewSum(num).GetSum(200)
	if sum != 2384 {
		t.Error("Expected: 2284, got:", sum)
	}

	t.Logf("Got MockThousandHundredNumber : %d", sum)
}
