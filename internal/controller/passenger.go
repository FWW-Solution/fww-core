package controller

import (
	"fmt"
	"fww-core/internal/data/dto"
	"fww-core/internal/data/dto_passanger"
	"fww-core/internal/tools"
	"log"
	"strconv"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) RegisterPassanger(ctx *fiber.Ctx) error {
	var body dto_passanger.RequestRegister

	if err := ctx.BodyParser(&body); err != nil {
		log.Println(err)
		err := tools.ResponseBadRequest(err)
		c.Log.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	// validate body
	errValidation := tools.ValidateVariable(body)
	if errValidation != nil {
		c.Log.Error(errValidation)
		return ctx.Status(400).JSON(errValidation)
	}

	result, err := c.UseCase.RegisterPassanger(&body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	meta := dto.MetaResponse{
		StatusCode: "201",
		IsSuccess:  true,
		Message:    "Success",
	}

	response := tools.ResponseBuilder(result, meta)

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *Controller) DetailPassanger(ctx *fiber.Ctx) error {
	data := ctx.Query("id")
	dataInt, err := strconv.Atoi(data)
	dataInt64 := int64(dataInt)

	if err != nil {
		err := tools.ResponseBadRequest(err)
		c.Log.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	result, err := c.UseCase.DetailPassanger(dataInt64)
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

	fmt.Println(response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *Controller) UpdatePassanger(ctx *fiber.Ctx) error {
	var body dto_passanger.RequestUpdate

	if err := ctx.BodyParser(&body); err != nil {
		err := tools.ResponseBadRequest(err)
		c.Log.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	// validate body
	errValidation := tools.ValidateVariable(body)
	if errValidation != nil {
		c.Log.Error(errValidation)
		return ctx.Status(400).JSON(errValidation)
	}

	result, err := c.UseCase.UpdatePassanger(&body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	meta := dto.MetaResponse{
		StatusCode: "201",
		IsSuccess:  true,
		Message:    "Success",
	}

	response := tools.ResponseBuilder(result, meta)

	return ctx.Status(fiber.StatusCreated).JSON(response)

}

func (c *Controller) UpdatePassangerByIDNumberHandler(msg *message.Message) error {
	var body dto_passanger.RequestUpdateBPM

	if err := json.Unmarshal(msg.Payload, &body); err != nil {
		msg.Ack()
		return err
	}

	err := c.UseCase.UpdatePassangerByIDNumber(&body)
	if err != nil {
		msg.Ack()
		c.Log.Error(err)
		return err
	}

	msg.Ack()

	return nil
}
