package subscriber

import "fmt"

type Customer struct {
	ID   string
	Name string
}

func NewCustomer(id, name string) Customer {
	return Customer{ID: id, Name: name}
}

func (customer Customer) ReactToPublisherMsg(msg string) {
	fmt.Printf("Customer %s (%s) received: %s\n", customer.ID, customer.Name, msg)
}

func (customer Customer) GetID() string {
	return customer.ID
}
