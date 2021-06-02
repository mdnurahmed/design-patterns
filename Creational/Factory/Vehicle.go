package factory

type Vehicle struct {
	typeofvehicle string
	wheels        int
	maxSpeedMiles int
}

func (v *Vehicle) setType(typeofvehicle string) {
	v.typeofvehicle = typeofvehicle
}

func (v *Vehicle) setWheels(wheels int) {
	v.wheels = wheels
}

func (v *Vehicle) setMaxSpeedMiles(speed int) {
	v.maxSpeedMiles = speed
}
