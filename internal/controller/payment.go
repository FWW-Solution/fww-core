package controller

import (
	"fww-core/internal/data/dto"
	"fww-core/internal/data/dto_payment"
	"fww-core/internal/tools"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) RequestPayment(msg *message.Message) error {
	var req dto_payment.Request
	paymentCodeID := msg.UUID

	if err := json.Unmarshal(msg.Payload, &req); err != nil {
		msg.Ack()
		return err
	}

	if err := c.UseCase.RequestPayment(&req, paymentCodeID); err != nil {
		msg.Ack()
		return err
	}

	msg.Ack()

	return nil

}

func (c *Controller) GetPaymentStatus(ctx *fiber.Ctx) error {
	codePayment := ctx.Query("code_payment", "")

	result, err := c.UseCase.GetDetailPayment(codePayment)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	meta := dto.MetaResponse{
		StatusCode: "200",
		IsSuccess:  true,
		Message:    "Success",
	}

	response := tools.ResponseBuilder(result, meta)

	return ctx.Status(fiber.StatusOK).JSON(response)
}
