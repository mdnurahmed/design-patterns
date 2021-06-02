package strategy

type sortableCollection struct {
	collection      []int
	sortingStrategy ISortingStrategy
}

func (s *sortableCollection) sort() {
	s.sortingStrategy.sort(s.collection)
}

func (s *sortableCollection) setSortingStrategy(sortingStrategy ISortingStrategy) {
	s.sortingStrategy = sortingStrategy
}
