package adapter

import (
	"fww-core/internal/data/dto_passanger"

	"github.com/ThreeDotsLabs/watermill/message"
)

type adapter struct {
	pub message.Publisher
	sub message.Subscriber
}

type Adapter interface {
	CheckPassangerInformations(data *dto_passanger.RequestBPM) error
	// Payment
	RequestPayment(data interface{})
	SendNotification(data interface{})
}
