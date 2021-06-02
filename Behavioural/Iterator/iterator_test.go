package iterator

import "testing"

func TestGetInstance(t *testing.T) {

	collection := []int{1, 2, 3, 4}
	iterator := newZigZagIterator(collection)

	v1, err1 := iterator.getNext()
	if v1 != collection[0] {
		t.Errorf("Got value %d , expected value to be %d", v1, collection[0])
	}
	if err1 != nil {
		t.Errorf("Got error %s, expected error to be nil", err1.Error())
	}
	v2, err2 := iterator.getNext()
	if v2 != collection[3] {
		t.Errorf("Got value %d , expected value to be %d", v2, collection[3])
	}
	if err2 != nil {
		t.Errorf("Got error %s, expected error to be nil", err2.Error())
	}
}
