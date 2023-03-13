package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
	"server/database"
	"server/database/model"
)

func init() {
	AddApiNeedAuthArray(RoutStruct{Path: "/user", Function: getAllUser, Method: Method_Get})
}

func getAllUser(c *fiber.Ctx) error {
	var userList []model.User
	database.GetDB().Preload(clause.Associations).Find(&userList)
	return c.JSON(userList)
}
