package iterator18

type Iterator interface {
	HasNext() bool
	Next()
	CurrentItem() interface{}
}

type ArrayInt []int

func (a ArrayInt) Iterator() Iterator {
	return &ArrayIntIterator{
		a,
		0,
	}
}

type ArrayIntIterator struct {
	arrayInt ArrayInt
	index    int
}

func (a *ArrayIntIterator) HasNext() bool {
	return a.index < len(a.arrayInt)-1
}

func (a *ArrayIntIterator) Next() {
	a.index++
}

func (a *ArrayIntIterator) CurrentItem() interface{} {
	return a.arrayInt[a.index]
}
