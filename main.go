package main

import (
	"assignment7_Observer/observer"
	"assignment7_Observer/subject"
)

func main() {
	center := subject.NewCarRentalCenter("Toyota Camry")

	regularCustomer := observer.NewCustomer("C001", "John")
	vipCustomer := observer.NewVIPCustomer("V001", "Sarah")

	center.Register(regularCustomer)
	center.Register(vipCustomer)

	center.UpdateAvailability()

	center.Deregister(regularCustomer)
	center.Deregister(vipCustomer)
	//center.UpdateAvailability()
}
