package messagestream

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type goChannel struct {
}

func NewGoChannel() MessageStream {
	return &goChannel{}
}

func (m *goChannel) NewSubscriber() (message.Subscriber, error) {
	subscribe := gochannel.NewGoChannel(gochannel.Config{}, watermill.NewStdLogger(stateLog, stateLog))

	return subscribe, nil
}

func (m *goChannel) NewPublisher() (message.Publisher, error) {
	publisher := gochannel.NewGoChannel(gochannel.Config{}, watermill.NewStdLogger(stateLog, stateLog))

	return publisher, nil
}
