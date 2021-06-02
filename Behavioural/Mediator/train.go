package mediator

import "fmt"

const (
	Waiting      = "Waiting"
	InPlatform   = "InPlatform"
	LeftPlatform = "LeftPlatform"
	UNKNOWN      = "UNKNOWN"
)

type train struct {
	mediator IStationManager
	state    string
	ID       int
}

func (t *train) requestArrival() {
	if t.mediator.requestArrival(t) {
		t.state = InPlatform
	} else {
		t.state = Waiting
	}
}

func (t *train) departure() error {
	if t.state == InPlatform {
		t.state = LeftPlatform
		t.mediator.notifyFree()
		return nil
	}
	return fmt.Errorf("train %d is not at the platform", t.ID)
}

func (t *train) permitArrival() {
	t.state = InPlatform
}

var UID int

func newTrain(stationManager IStationManager) Itrain {
	UID++
	return &train{
		mediator: stationManager,
		ID:       UID,
		state:    "UNKNOWN",
	}
}
