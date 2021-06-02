package mediator

type IStationManager interface {
	requestArrival(Itrain) bool
	notifyFree()
}
