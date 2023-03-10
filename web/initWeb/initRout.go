package initWeb

import (
	"github.com/gofiber/fiber/v2"
	"server/config"
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
	c.BodyParser(&bodyRequest)
	ServerUser := model.User{Name: "Server Init User", Username: "server"}
	ServerUser.Save()
	ServerUser.SetCreatedByAndSave(&ServerUser)

	bodyRequest.ServerInit.Save()
	bodyRequest.ServerInit.SetCreatedByAndSave(&ServerUser)

	bodyRequest.UserInit.SetPasswordWithBcrypt(bodyRequest.UserInit.Password)
	bodyRequest.UserInit.Save()
	bodyRequest.UserInit.SetCreatedByAndSave(&ServerUser)

	AdminRole := model.RoleUser{Name: "Administrator", Code: "moderator", UserContain: []*model.User{&ServerUser, &bodyRequest.UserInit}}
	AdminRole.Save()
	AdminRole.SetCreatedByAndSave(&ServerUser)

	config.GetAppProperties().NeedSetup = false
	return c.JSON(bodyRequest)
}

type setupRequestBody struct {
	UserInit   model.User
	ServerInit model.ServerDetails
}
