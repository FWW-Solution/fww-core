package adapter

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type adapter struct {
	pub *message.Publisher
	sub *message.Subscriber
}

type Adapter interface {
	CheckPassangerInformations(data interface{})
	// Payment
	RequestPayment(data interface{})
	SendNotification(data interface{})
}
