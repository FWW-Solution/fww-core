package controller

import (
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	useCase interface{}
}

func (c *Controller) RegisterPassanger(ctx *fiber.Ctx) error {
	return nil
}

func (c *Controller) DetailPassanger(ctx *fiber.Ctx) error {
	return nil
}

func (c *Controller) UpdatePassanger(ctx *fiber.Ctx) error {
	return nil
}
