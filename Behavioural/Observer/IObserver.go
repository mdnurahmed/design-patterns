package observer

type IObserver interface {
	notify(string)
	getID() string
}
