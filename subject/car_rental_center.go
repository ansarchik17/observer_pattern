package subject

import "assignment7_Observer/customer_interface"

type CarRentalCenter struct {
	observers []customer_interface.Customer
	carName   string
	available bool
}

func NewCarRentalCenter(carName string) *CarRentalCenter {
	return &CarRentalCenter{carName: carName}
}
