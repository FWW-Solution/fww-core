package adapter

import (
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/data/dto_passanger"
	"fww-core/internal/data/dto_payment"
	"fww-core/internal/data/dto_ticket"

	"github.com/ThreeDotsLabs/watermill/message"
)

type adapter struct {
	pub message.Publisher
	sub message.Subscriber
}



type Adapter interface {
	// Passanger
	CheckPassangerInformations(data *dto_passanger.RequestBPM) error
	// Payment
	RequestGenerateInvoice(data *dto_booking.RequestBPM) error
	DoPayment(data *dto_payment.DoPayment) error
	RedeemTicket(data *dto_ticket.RequestRedeemTicketToBPM) error
	SendNotification(data interface{})
}
