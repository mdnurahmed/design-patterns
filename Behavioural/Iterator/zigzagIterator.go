package iterator

import "fmt"

type zigzagIterator struct {
	returnLeftIndex bool
	leftindex       int
	rightindex      int
	collection      []int
}

func (z *zigzagIterator) hasNext() bool {
	return z.leftindex < z.rightindex {
}

func (z *zigzagIterator) getNext() (int, error) {
	var res int
	if z.hasNext() {
		if z.returnLeftIndex {
			res = z.collection[z.leftindex]
			z.leftindex++
			z.returnLeftIndex = false
		} else {
			res = z.collection[z.rightindex]
			z.rightindex++
			z.returnLeftIndex = true
		}
		return res, nil
	}
	return res, fmt.Errorf("Has no element left")
}

func newZigZagIterator(collection []int) IZigzagIterator {
	return &zigzagIterator{
		collection:      collection,
		returnLeftIndex: true,
		leftindex:       0,
		rightindex:      len(collection) - 1,
	}
}
