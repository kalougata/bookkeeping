package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/bookkeeping/internal/controller"
)

func NewHTTPServer(
	authC *controller.AuthController,
) *fiber.App {
	app := fiber.New()

	app.Get("", func(ctx *fiber.Ctx) error {
		return ctx.SendString("HelloWorld")
	})

	return app
}