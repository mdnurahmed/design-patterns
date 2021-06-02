package factory

type Car struct {
	Vehicle
}

func newCar(typeofvehicle string, maxSpeedMiles, wheels int) IVehicle {
	return &Car{
		Vehicle: Vehicle{
			typeofvehicle: typeofvehicle,
			maxSpeedMiles: maxSpeedMiles,
			wheels:        wheels,
		},
	}
}
