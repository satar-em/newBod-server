package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"server/web/jwt"
)

func init() {
	AddRoutArray(RoutInit{Path: "/hello", Function: GetHello, Metod: Method_Get})
}

func GetHello(c *fiber.Ctx) error {
	if c.GetReqHeaders()["Authentication"] != "" {
		auth := jwt.GetAuth(c.GetReqHeaders()["Authentication"])
		return c.SendString(fmt.Sprintf("Hello your user id is %d", *auth.GetUser()))
	}
	return c.SendString("Hello,world")
}
