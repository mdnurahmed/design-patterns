package factory

type Cycle struct {
	Vehicle
}

func newCycle(typeofvehicle string, maxSpeedMiles, wheels int) IVehicle {
	return &Cycle{
		Vehicle: Vehicle{
			typeofvehicle: typeofvehicle,
			maxSpeedMiles: maxSpeedMiles,
			wheels:        wheels,
		},
	}
}
