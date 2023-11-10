package controller

import (
	"fww-core/internal/data/dto"
	"fww-core/internal/tools"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) ListAirport(ctx *fiber.Ctx) error {
	city := ctx.Query("city")
	province := ctx.Query("province")
	iata := ctx.Query("iata")

	result, err := c.UseCase.GetAirport(city, province, iata)
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
