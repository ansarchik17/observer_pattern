package publisher

import (
	"assignment7_Observer/subscriber"
	"fmt"
)

type CarRentalOffice struct {
	name        string
	subscribers []subscriber.Subscriber
}

func NewCarRentalOffice(name string) *CarRentalOffice {
	return &CarRentalOffice{
		name:        name,
		subscribers: make([]subscriber.Subscriber, 0),
	}
}

func (office *CarRentalOffice) Register(subscriber subscriber.Subscriber) {
	office.subscribers = append(office.subscribers, subscriber)
}

func (office *CarRentalOffice) NotifyAll(message string) {
	fmt.Printf("\n [%s] Sending notification: \"%s\"\n", office.name, message)
	for _, subscriber := range office.subscribers {
		fmt.Println("Notifying:", subscriber.GetID())
		subscriber.ReactToPublisherMsg(message)
	}
}
