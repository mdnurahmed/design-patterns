package observer

type item struct {
	observerList map[IObserver]struct{}
	name         string
	inStock      bool
}

func (i *item) updateAvailability(isAvailable bool) {
	i.inStock = isAvailable
	if i.inStock {
		i.notifyAllObservers()
	}
}

func (i *item) register(o IObserver) {
	i.observerList[o] = struct{}{}
}

func (i *item) deregister(o IObserver) {
	_, ok := i.observerList[o]
	if ok {
		delete(i.observerList, o)
	}
}

func (i *item) notifyAllObservers() {
	for observer := range i.observerList {
		observer.notify(i.name)
	}
}

func newItem(name string) *item {
	return &item{
		name:         name,
		observerList: make(map[IObserver]struct{}),
	}
}
