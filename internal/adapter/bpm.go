package adapter

import (
	"fmt"
	"fww-core/internal/data/dto_passanger"

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

// RequestPayment implements Adapter.
func (a *adapter) RequestPayment(data interface{}) {
	fmt.Println("Do Payment")
}

// SendNotification implements Adapter.
func (a *adapter) SendNotification(data interface{}) {
	fmt.Println("Send Notification")
}
