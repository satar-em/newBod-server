package router

import (
	"github.com/gofiber/fiber/v2"
	"server/database/repo/user"
)

func PostLogin(c *fiber.Ctx) error {
	var login LoginBody
	err := c.BodyParser(&login)
	if err != nil {
		return c.SendString(err.Error())
	}
	userLogin := user.GetUserByUserAndPass(login.Username, login.Password)
	if userLogin != nil {
		return c.JSON(userLogin)
	}
	return c.SendString("Not Found User")
}

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
