package controller

import (
	"errors"
	"fww-core/internal/data/dto"
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/tools"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
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
		log.Println(err)
		msg.Ack()
		return err
	}
	return nil
}

func (c *Controller) GetDetailBooking(ctx *fiber.Ctx) error {
	codeBooking := ctx.Query("code_booking", "")

	result, err := c.UseCase.GetDetailBooking(codeBooking)
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
