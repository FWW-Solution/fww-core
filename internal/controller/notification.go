package controller

import (
	"fww-core/internal/data/dto_notification"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
)

func (c *Controller) SendEmailNotificationHandler(msg *message.Message) error {
	var body dto_notification.Request

	if err := json.Unmarshal(msg.Payload, &body); err != nil {
		msg.Ack()
		c.Log.Error(err)
		return err
	}

	err := c.UseCase.InquiryNotification(&body)
	if err != nil {
		msg.Ack()
		c.Log.Error(err)
		return err
	}

	msg.Ack()
	return nil
}
