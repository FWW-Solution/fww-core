package adapter

import (
	"fmt"
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/data/dto_passanger"
	"fww-core/internal/data/dto_payment"
	"fww-core/internal/data/dto_ticket"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
)

func NewBPM(publisher message.Publisher, subscriber message.Subscriber) Adapter {
	return &adapter{
		pub: publisher,
		sub: subscriber,
	}
}

// CheckPassangerInformations implements Adapter.
func (a *adapter) CheckPassangerInformations(data *dto_passanger.RequestBPM) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err

	}

	ID := watermill.NewUUID()
	err = a.pub.Publish("start_process_passanger", message.NewMessage(ID, json))
	if err != nil {
		return err
	}

	return nil
}

// RequestGenerateInvoice implements Adapter.
func (a *adapter) RequestGenerateInvoice(data *dto_booking.RequestBPM) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ID := watermill.NewUUID()
	err = a.pub.Publish("start_process_booking", message.NewMessage(ID, json))
	if err != nil {
		return err
	}

	return nil
}

// RequestPayment implements Adapter.
func (a *adapter) DoPayment(data *dto_payment.DoPayment) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ID := watermill.NewUUID()
	err = a.pub.Publish("do_payment_bpm", message.NewMessage(ID, json))
	if err != nil {
		return err
	}

	return nil
}

// RedeemTicket implements Adapter.
func (a *adapter) RedeemTicket(data *dto_ticket.RequestRedeemTicketToBPM) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ID := watermill.NewUUID()
	err = a.pub.Publish("redeem_ticket_bpm", message.NewMessage(ID, json))
	if err != nil {
		return err
	}

	return nil
}

// SendNotification implements Adapter.
func (a *adapter) SendNotification(data interface{}) {
	fmt.Println("Send Notification")
}
