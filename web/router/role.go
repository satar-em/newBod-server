package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
	"server/database"
	"server/database/model"
)

func init() {
	AddApiNeedAuthArray(RoutStruct{Path: "/role", Function: getAllRole, Method: Method_Get})
}

func getAllRole(c *fiber.Ctx) error {
	var roleList []model.RoleUser
	database.GetDB().Preload(clause.Associations).Find(&roleList)
	return c.JSON(roleList)
}
