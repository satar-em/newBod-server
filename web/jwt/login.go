package jwt

import (
	"github.com/gofiber/fiber/v2"
	"server/config"
	"server/database/repo/user"
)

func PostLogin(c *fiber.Ctx) error {
	loginBody := map[string]string{}
	err := c.BodyParser(&loginBody)
	if err != nil {
		return c.SendString(err.Error())
	}
	username := loginBody[config.GetAppProperties().WebServer.Security.UsernameParamName]
	password := loginBody[config.GetAppProperties().WebServer.Security.PasswordParamName]
	userLogin := user.GetUserByUserAndPass(username, password)
	if userLogin != nil {
		authentication := GetAuthWithUser(userLogin.ID)
		authentication.Fresh()
		return c.JSON(LoginSuccessResponse{Token: *authentication.GetId(), User: userLogin.Name})
	}
	return c.JSON(LoginErrorResponse{Message: "cannot login bro :)"})
}

type LoginErrorResponse struct {
	Message string `json:"message"`
}
type LoginSuccessResponse struct {
	Token string `json:"token"`
	User  string `json:"user"`
}
