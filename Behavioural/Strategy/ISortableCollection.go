package strategy

type ISortableCollection interface {
	sort()
	setSortingStrategy(ISortingStrategy)
}
