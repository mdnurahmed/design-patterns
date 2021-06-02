package factory

type IVehicle interface {
	setType(typeofvehicle string)
	setWheels(val int)
	setMaxSpeedMiles(val int)
}
