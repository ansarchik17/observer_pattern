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

func (o *CarRentalOffice) Register(sub subscriber.Subscriber) {
	o.subscribers = append(o.subscribers, sub)
}

func (o *CarRentalOffice) NotifyAll(msg string) {
	fmt.Printf("\n [%s] Sending notification: \"%s\"\n", o.name, msg)
	for _, sub := range o.subscribers {
		fmt.Println("Notifying:", sub.GetID())
		sub.ReactToPublisherMsg(msg)
	}
}
