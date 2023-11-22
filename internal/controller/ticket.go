package controller

import (
	"fww-core/internal/data/dto"
	"fww-core/internal/data/dto_ticket"
	"fww-core/internal/tools"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) RedeemTicket(ctx *fiber.Ctx) error {
	var body dto_ticket.Request

	if err := ctx.BodyParser(&body); err != nil {
		c.Log.Error(err)
		response := tools.ResponseBadRequest(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := c.UseCase.RedeemTicket(body.CodeBooking)
	if err != nil {
		c.Log.Error(err)
		response := tools.ResponseInternalServerError(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	meta := dto.MetaResponse{
		StatusCode: "OK200",
		IsSuccess:  true,
		Message:    "Ticket redeemed",
	}

	response := tools.ResponseBuilder(result, meta)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *Controller) UpdateTicketHandler(msg *message.Message) error {
	var body dto_ticket.RequestUpdateTicket

	if err := json.Unmarshal(msg.Payload, &body); err != nil {
		msg.Ack()
		c.Log.Error(err)
		return err
	}
	err := c.UseCase.UpdateTicket(&body)
	if err != nil {
		msg.Ack()
		c.Log.Error(err)
		return err
	}
	msg.Ack()
	return nil
}
