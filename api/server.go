package api

import (
	"polar/core"

	"github.com/gofiber/fiber/v2"
)

func Listen() {
	app := fiber.New()

	app.Post("/push/:uid", func(c *fiber.Ctx) error {
		return core.Receive(core.Request{})
	})

	app.Get("/pull/:uid", func(c *fiber.Ctx) error {
		return core.Send(core.Request{})
	})

	app.Delete("/erase/:uid", func(c *fiber.Ctx) error {
		return core.Erase(core.Request{})
	})

	app.Listen(":3000")
}
