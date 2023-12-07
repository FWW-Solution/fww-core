package controller

import (
	"fww-core/internal/data/dto"
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/tools"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// func (c *Controller) RequestBooking(msg *message.Message) error {
// 	var body dto_booking.Request
// 	bookingIDCode := msg.UUID
// 	err := json.Unmarshal(msg.Payload, &body)
// 	if err != nil {
// 		c.Log.Error(err)
// 		return err
// 	}

// 	// validate body
// 	errValidation := tools.ValidateVariable(body)
// 	if errValidation != nil {
// 		c.Log.Error(errValidation)
// 		msg.Ack()
// 		return errors.New("validation error")
// 	}

// 	err = c.UseCase.RequestBooking(&body, bookingIDCode)
// 	if err != nil {
// 		c.Log.Error(err)
// 		msg.Ack()
// 		return err
// 	}
// 	msg.Ack()
// 	return nil
// }

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

func (c *Controller) UpdateBookingHandler(msg *message.Message) error {
	var body dto_booking.RequestUpdateBooking

	if err := json.Unmarshal(msg.Payload, &body); err != nil {
		msg.Ack()
		c.Log.Error(err)
		return err
	}
	err := c.UseCase.UpdateBooking(&body)
	if err != nil {
		msg.Ack()
		c.Log.Error(err)
		return err
	}
	msg.Ack()
	return nil
}

func (c *Controller) RequestBooking(ctx *fiber.Ctx) error {
	var body dto_booking.Request

	uuid := watermill.NewUUID()

	if err := ctx.BodyParser(&body); err != nil {
		err := tools.ResponseBadRequest(err)
		c.Log.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := c.UseCase.RequestBooking(&body, uuid)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	meta := dto.MetaResponse{
		StatusCode: "200",
		IsSuccess:  true,
		Message:    "Success",
	}

	response := tools.ResponseBuilder(nil, meta)

	return ctx.Status(fiber.StatusOK).JSON(response)
}
