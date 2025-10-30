package customer_interface

type Customer interface {
	Update(carName string)
	GetID() string
}
