package subscriber

type Subscriber interface {
	ReactToPublisherMsg(msg string)
	GetID() string
}
