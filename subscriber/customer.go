package subscriber

import "fmt"

type Customer struct {
	ID   string
	Name string
}

func NewCustomer(id, name string) Customer {
	return Customer{ID: id, Name: name}
}

func (c Customer) ReactToPublisherMsg(msg string) {
	fmt.Printf("Customer %s (%s) received: %s\n", c.ID, c.Name, msg)
}

func (c Customer) GetID() string {
	return c.ID
}
