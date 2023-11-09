package adapter

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

func NewBPM(publisher *message.Publisher, subscriber *message.Subscriber) Adapter {
	return &adapter{
		pub: publisher,
		sub: subscriber,
	}
}

// CheckPassangerInformations implements Adapter.
func (a *adapter) CheckPassangerInformations(data interface{}) {
	panic("unimplemented")
}
