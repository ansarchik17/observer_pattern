package observer

import (
	"fmt"
)

type RegularCustomer struct {
	id   string
	name string
}

func NewCustomer(id, name string) *RegularCustomer {
	return &RegularCustomer{id: id, name: name}
}

func (regularCustomer *RegularCustomer) Update(carName string) {
	fmt.Printf("1)Regular Customer %s: '%s' is now available for rent!\n", regularCustomer.name, carName)
}

func (regularCustomer *RegularCustomer) GetID() string {
	return regularCustomer.id
}
