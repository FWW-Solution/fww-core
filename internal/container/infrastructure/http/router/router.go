package router

import (
	"fww-core/internal/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Initialize(app *fiber.App, ctrl *controller.Controller) *fiber.App {
	app.Get("/", monitor.New(monitor.Config{Title: "fww-core metrics page"}))

	Api := app.Group("/api")

	v1 := Api.Group("/private/v1")

	v1.Post("/passanger", ctrl.RegisterPassanger)
	v1.Get("/passanger", ctrl.DetailPassanger)
	v1.Put("/passanger", ctrl.UpdatePassanger)

	return app

}
