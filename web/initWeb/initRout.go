package initWeb

import (
	"github.com/gofiber/fiber/v2"
	"server/config"
)

func initStartForFistTime(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		if c.Path() == config.GetAppProperties().WebServer.SetupPath {
			if config.GetAppProperties().NeedSetup {
				return c.Next()
			}
			return c.SendString("canSetup = false")
		}
		if config.GetAppProperties().NeedSetup {
			return c.Redirect(config.GetAppProperties().WebServer.SetupPath)
		}
		return c.Next()
	})
	app.Get(config.GetAppProperties().WebServer.SetupPath, func(c *fiber.Ctx) error {
		return c.SendString("canSetup = true")
	})
}
