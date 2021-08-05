package sum

type Sum struct {
	meth SumX
}

type SumX interface {
	GetNumber() int
}

func NewSum(meth SumX) *Sum {
	return &Sum{meth}
}

func (s *Sum) GetSum(v int) int {
	n := v + s.meth.GetNumber()
	return n
}
