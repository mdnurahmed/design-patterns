package mediator

import "testing"

func TestMediator(t *testing.T) {
	stationManager := newStationManger()
	train1 := newTrain(stationManager).(*train)
	train2 := newTrain(stationManager).(*train)
	train1.requestArrival()
	if train1.state != InPlatform {
		t.Errorf("expected state to be %s but got %s \n", InPlatform, train1.state)
	}
	train2.requestArrival()
	if train2.state != Waiting {
		t.Errorf("expected state to be %s but got %s \n", Waiting, train2.state)
	}
	err := train1.departure()
	if err != nil {
		t.Errorf("expected error to be nil but got %s \n", err.Error())
	}
	if train1.state != LeftPlatform {
		t.Errorf("expected state to be %s but got %s \n", LeftPlatform, train1.state)
	}
	if train2.state != InPlatform {
		t.Errorf("expected state to be %s but got %s \n", InPlatform, train2.state)
	}
}
