package sum

type Sum struct {
	GetThousandHundred func() int
}

func (s *Sum) Sum(v int) int {
	n := v + s.GetThousandHundred()
	return n
}
