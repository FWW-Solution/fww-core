package controller

import (
	"fww-core/internal/data/dto"
	"fww-core/internal/data/dto_passanger"
	"fww-core/internal/tools"
	"fww-core/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Controller struct {
	UseCase usecase.UseCase
	Log     *zap.SugaredLogger
}

func (c *Controller) RegisterPassanger(ctx *fiber.Ctx) error {
	var body dto_passanger.RequestRegister

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

	return ctx.Status(fiber.StatusAccepted).JSON(response)
}

func (c *Controller) DetailPassanger(ctx *fiber.Ctx) error {
	data := ctx.Query("data")
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
		StatusCode: "201",
		IsSuccess:  true,
		Message:    "Success",
	}

	response := tools.ResponseBuilder(result, meta)

	return ctx.Status(fiber.StatusAccepted).JSON(response)
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
