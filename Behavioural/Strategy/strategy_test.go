package strategy

import "testing"

func TestGetInstance(t *testing.T) {
	collection1 := sortableCollection{
		collection:      []int{3, 1, 2},
		sortingStrategy: new(BubbleSortStrategy),
	}
	collection1.sort()
	for i := 0; i < len(collection1.collection); i++ {
		if collection1.collection[i] != i+1 {
			t.Errorf("Collection is not sorted. Expected %d but got %d\n", i+1, collection1.collection[i])
		}
	}
	collection2 := sortableCollection{
		collection:      []int{3, 1, 2},
		sortingStrategy: new(CountingSortStrategy),
	}
	collection2.sort()
	for i := 0; i < len(collection2.collection); i++ {
		if collection2.collection[i] != i+1 {
			t.Errorf("Collection is not sorted. Expected %d but got %d\n", i+1, collection2.collection[i])
		}
	}

}
