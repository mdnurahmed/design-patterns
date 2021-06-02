package factory

import "testing"

func TestVehicleFactory(t *testing.T) {
	vehicle1, err1 := getVehicle("car")
	car := vehicle1.(*Car)
	if car.typeofvehicle != "car" {
		t.Errorf("Got %s, expected typeofvehicle to be car\n", car.typeofvehicle)
	}
	if car.wheels != 4 {
		t.Errorf("Got %d, expected no of wheels to be 4\n", car.wheels)
	}
	if car.maxSpeedMiles != 250 {
		t.Errorf("Got %d, expected maxSpeedMiles to be 250\n", car.maxSpeedMiles)
	}
	if err1 != nil {
		t.Errorf("Got error %s, expected error to be nil\n", err1.Error())
	}

	vehicle2, err2 := getVehicle("bike")
	bike := vehicle2.(*Bike)
	if bike.typeofvehicle != "bike" {
		t.Errorf("Got %s, expected typeofvehicle to be bike\n", bike.typeofvehicle)
	}
	if bike.wheels != 2 {
		t.Errorf("Got %d, expected no of wheels to be 2\n", bike.wheels)
	}
	if bike.maxSpeedMiles != 175 {
		t.Errorf("Got %d, expected maxSpeedMiles to be 250\n", bike.maxSpeedMiles)
	}
	if err2 != nil {
		t.Errorf("Got error %s, expected error to be nil\n", err2.Error())
	}

	vehicle3, err3 := getVehicle("cycle")
	cycle := vehicle3.(*Cycle)
	if cycle.typeofvehicle != "cycle" {
		t.Errorf("Got %s, expected typeofvehicle to be cycle\n", cycle.typeofvehicle)
	}
	if cycle.wheels != 2 {
		t.Errorf("Got %d, expected no of wheels to be 2\n", cycle.wheels)
	}
	if cycle.maxSpeedMiles != 60 {
		t.Errorf("Got %d, expected no of wheels to be 60\n", cycle.maxSpeedMiles)
	}
	if err3 != nil {
		t.Errorf("Got error %s, expected error to be nil\n", err3.Error())
	}
	vehicle4, err4 := getVehicle("horse")
	if vehicle4 != nil {
		t.Errorf("Got %+v, expected vehicle to be nil\n", vehicle4)
	}
	if err4 == nil {
		t.Errorf("Got nil, expected error\n")
	}
}
