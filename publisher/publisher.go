package publisher

import "assignment7_Observer/subscriber"

type Publisher interface {
	Register(sub subscriber.Subscriber)
	NotifyAll(msg string)
}
