package router

import "github.com/gofiber/fiber/v2"

type RoutInit struct {
	Path     string
	Function func(c *fiber.Ctx) error
	Metod    int
}

const (
	Method_Get    = 0
	Method_Post   = 1
	Method_Put    = 2
	Method_Delete = 3
)

var routerArray []RoutInit

func AddRoutArray(rout RoutInit) {
	routerArray = append(routerArray, rout)
}

func GetRoutArray() []RoutInit {
	return routerArray
}
