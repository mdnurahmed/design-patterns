package observer

type ISubject interface {
	register(Observer IObserver)
	deregister(Observer IObserver)
	notifyAllObserver()
}
