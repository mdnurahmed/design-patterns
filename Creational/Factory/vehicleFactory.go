package factory

import "fmt"

func getVehicle(vehicleType string) (IVehicle, error) {
	if vehicleType == "car" {
		return newCar("car", 250, 4), nil
	}
	if vehicleType == "bike" {
		return newBike("bike", 175, 2), nil
	}
	if vehicleType == "cycle" {
		return newCycle("cycle", 60, 2), nil
	}
	return nil, fmt.Errorf("Wrong vehicle type passed")
}
