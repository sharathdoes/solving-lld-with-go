package internal

type VehicleType string

const (
	VehicleTypeCar   VehicleType = "Car"
	VehicleTypeBike  VehicleType = "Bike"
	VehicleTypeTruck VehicleType = "Truck"
)

type Vehicle struct {
	VehicleNumber string
	Type VehicleType
}
