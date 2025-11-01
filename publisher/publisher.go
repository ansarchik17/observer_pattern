package publisher

import "assignment7_Observer/subscriber"

type Publisher interface {
	Register(subscriber subscriber.Subscriber)
	NotifyAll(msg string)
}
