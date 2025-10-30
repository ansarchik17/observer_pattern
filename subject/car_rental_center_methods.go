package subject

import (
	"assignment7_Observer/customer_interface"
	"fmt"
)

func (carrentalcenter *CarRentalCenter) Register(o customer_interface.Customer) {
	carrentalcenter.observers = append(carrentalcenter.observers, o)
}

func (carrentalcenter *CarRentalCenter) Deregister(o customer_interface.Customer) {
	length := len(carrentalcenter.observers)
	for i, obs := range carrentalcenter.observers {
		if obs.GetID() == o.GetID() {
			carrentalcenter.observers[i], carrentalcenter.observers[length-1] = carrentalcenter.observers[length-1], carrentalcenter.observers[i]
			carrentalcenter.observers = carrentalcenter.observers[:length-1]
			return
		}
	}
}

func (carrentalcenter *CarRentalCenter) NotifyAll() {
	for _, observer := range carrentalcenter.observers {
		observer.Update(carrentalcenter.carName)
	}
}

func (carrentalcenter *CarRentalCenter) UpdateAvailability() {
	fmt.Printf("Car '%s' is now available for rent!\n", carrentalcenter.carName)
	carrentalcenter.available = true
	carrentalcenter.NotifyAll()
}
