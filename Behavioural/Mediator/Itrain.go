package mediator

type Itrain interface {
	requestArrival()
	departure() error
	permitArrival()
}
