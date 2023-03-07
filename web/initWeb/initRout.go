package initWeb

import (
	"github.com/gofiber/fiber/v2"
	"server/config"
	"server/database"
	"server/database/model"
)

func initStartForFistTime(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		if c.Path() == config.GetAppProperties().WebServer.SetupPath {
			if config.GetAppProperties().NeedSetup {
				return c.Next()
			}
			return c.Redirect(config.GetAppProperties().WebServer.SetupPath + "/error")
		}
		if config.GetAppProperties().NeedSetup {
			return c.Redirect(config.GetAppProperties().WebServer.SetupPath)
		}
		return c.Next()
	})
	app.Get(config.GetAppProperties().WebServer.SetupPath, getSetup)
	app.Post(config.GetAppProperties().WebServer.SetupPath, postSetup)
	app.Get(config.GetAppProperties().WebServer.SetupPath+"/error", getCantSetup)
}
func getCantSetup(c *fiber.Ctx) error {
	mapResponse := map[string]any{"canSetup": false}
	return c.JSON(mapResponse)
}
func getSetup(c *fiber.Ctx) error {
	mapResponse := map[string]any{"canSetup": true}
	return c.JSON(mapResponse)
}
func postSetup(c *fiber.Ctx) error {
	bodyRequest := setupRequestBody{}
	err := c.BodyParser(&bodyRequest)
	if err != nil {
		return c.SendString("bad Body")
	}
	database.GetDB().Create(&bodyRequest.UserInit)
	if bodyRequest.UserInit.ID != 0 {
		config.GetAppProperties().NeedSetup = false
	}
	return c.JSON(bodyRequest)
}

type setupRequestBody struct {
	UserInit model.User
}
