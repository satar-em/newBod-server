package router

import "github.com/gofiber/fiber/v2"

func init() {
	AddRoutArray(RoutInit{Path: "/hello", Function: GetHello, Metod: Method_Get})
}

func GetHello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
