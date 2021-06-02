package factory

type Bike struct {
	Vehicle
}

func newBike(typeofvehicle string, maxSpeedMiles, wheels int) IVehicle {
	return &Bike{
		Vehicle: Vehicle{
			typeofvehicle: typeofvehicle,
			maxSpeedMiles: maxSpeedMiles,
			wheels:        wheels,
		},
	}
}
