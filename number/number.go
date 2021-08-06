package number

type Number struct {
	number int
	method NumberMethod
}

type NumberMethod interface {
	SetRandomNumber(int64)
	GetRandomNumber() int
}

func NewNumber(method NumberMethod) *Number {
	return &Number{method: method}
}

func (n *Number) SetNumber(h int64) {
	n.method.SetRandomNumber(h)
}

func (n *Number) LoadNumber() {
	n.number = n.method.GetRandomNumber()
}

func (n *Number) GetNumber() int {
	return n.number
}
