package router

import "github.com/gofiber/fiber/v2"

type RoutStruct struct {
	Path     string
	Function func(c *fiber.Ctx) error
	Method   int
}

const (
	Method_Get    = 0
	Method_Post   = 1
	Method_Put    = 2
	Method_Delete = 3
)

var publicRouterArray []RoutStruct

func AddPublicRouterArray(rout RoutStruct) {
	publicRouterArray = append(publicRouterArray, rout)
}

func GetPublicRouterArray() []RoutStruct {
	return publicRouterArray
}

var apiNeedAuthArray []RoutStruct

func AddApiNeedAuthArray(rout RoutStruct) {
	apiNeedAuthArray = append(apiNeedAuthArray, rout)
}

func GetApiNeedAuthArray() []RoutStruct {
	return apiNeedAuthArray
}
