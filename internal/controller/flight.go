package controller

import (
	"fww-core/internal/data/dto"
	"fww-core/internal/tools"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetFlights(ctx *fiber.Ctx) error {
	departureTime := ctx.Query("departure_time", "")
	ArrivalTime := ctx.Query("arrivale_time", "")
	limit := ctx.Query("limit", "10")
	offset := ctx.Query("offset", "0")
	limitInt, err := tools.StringToInt(limit)
	if err != nil {
		c.Log.Error(err)
		return err
	}
	offsetInt, err := tools.StringToInt(offset)
	if err != nil {
		c.Log.Error(err)
		return err
	}
	result, err := c.UseCase.GetFlights(departureTime, ArrivalTime, limitInt, offsetInt)
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

func (c *Controller) GetDetailFlightByID(ctx *fiber.Ctx) error {
	id := ctx.Query("id", "")

	idInt64, err := tools.StringToInt64(id)

	result, err := c.UseCase.GetDetailFlightByID(idInt64)
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
