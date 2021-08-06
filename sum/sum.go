package sum

type Sum struct {
	method SumMethod
}

type SumMethod interface {
	GetNumber() int
}

func NewSum(method SumMethod) *Sum {
	return &Sum{method}
}

func (s *Sum) GetSum(n int) int {
	return n + s.method.GetNumber()
}
