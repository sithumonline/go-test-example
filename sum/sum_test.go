package sum

import (
	"testing"

	"github.com/sithumonline/go-test-example/number"
)

func MockHundredNumber() int {
	n := 184
	return n
}

func MockThousandHundredNumber() int {
	n := 2184
	return n
}

func TestSum(t *testing.T) {
	y := Sum{
		GetThousandHundred: number.Number{
			GetHundred: MockHundredNumber,
		}.GetHundred,
	}

	m := y.Sum(100)

	if m != 284 {
		t.Error("Expected: 284, got:", m)
	}

	t.Logf("Got MockHundredNumber : %d", m)

	z := Sum{
		GetThousandHundred: MockThousandHundredNumber,
	}

	u := z.Sum(200)

	if u != 2384 {
		t.Error("Expected: 2284, got:", u)
	}

	t.Logf("Got MockThousandHundredNumber : %d", u)
}
