package observer

import (
	"fmt"
)

type VIPCustomer struct {
	id   string
	name string
}

func NewVIPCustomer(id, name string) *VIPCustomer {
	return &VIPCustomer{id: id, name: name}
}

func (vipCustomer *VIPCustomer) Update(carName string) {
	fmt.Printf("2)VIP Customer %s: Priority notification! '%s' is available now!\n", vipCustomer.name, carName)
}

func (vipCustomer *VIPCustomer) GetID() string {
	return vipCustomer.id
}
