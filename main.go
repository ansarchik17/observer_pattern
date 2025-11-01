package main

import (
	"assignment7_Observer/publisher"
	"assignment7_Observer/subscriber"
)

func main() {
	office := publisher.NewCarRentalOffice("Astana Rental Car")

	customer1 := subscriber.NewCustomer("1", "Ansar")
	customer2 := subscriber.NewCustomer("2", "Aruzhan")
	customer3 := subscriber.NewCustomer("3", "Dias")
	customer4 := subscriber.NewCustomer("4", "Asylbek")

	office.Register(customer1)
	office.Register(customer2)
	office.Register(customer3)
	office.Register(customer4)

	office.NotifyAll("!!!New Lexus RX 2024 is now available!")
	office.NotifyAll("!!!Weekend discount: 20% off SUVs!")
}
