package iterator

type IZigzagIterator interface {
	hasNext() bool
	getNext() (int, error)
}
