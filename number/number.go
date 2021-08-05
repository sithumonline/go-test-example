package number

type Number struct {
	number int
	meth   NumberX
}

type NumberX interface {
	SetRandomNumber(int64)
	GetRandomNumber() int
}

func NewNumber(meth NumberX) *Number {
	return &Number{meth: meth}
}

func (n *Number) SetNumber(h int64) {
	n.meth.SetRandomNumber(h)
}

func (n *Number) LoadNumber() {
	n.number = n.meth.GetRandomNumber()
}

func (n *Number) GetNumber() int {
	return n.number
}
