package number

type Number struct {
	GetHundred func() int
}

func (n *Number) ThousandHundredNumber() int {
	q := 1000 + n.GetHundred()
	return q
}
