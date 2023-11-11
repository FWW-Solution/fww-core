package controller

import (
	"errors"
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/tools"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
)

func (c *Controller) RequestBooking(msg *message.Message) error {
	var body dto_booking.Request
	bookingIDCode := msg.UUID
	err := json.Unmarshal(msg.Payload, &body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	// validate body
	errValidation := tools.ValidateVariable(body)
	if errValidation != nil {
		c.Log.Error(errValidation)
		msg.Ack()
		return errors.New("validation error")
	}

	err = c.UseCase.RequestBooking(&body, bookingIDCode)
	if err != nil {
		c.Log.Error(err)
		msg.Ack()
		return err
	}
	return nil

}
