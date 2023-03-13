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
	/*bytes, err := json.Marshal(userList)
	if err != nil {
		log.Println(err)
	}
	var x []map[string]interface{}
	err = json.Unmarshal(bytes, &x)
	if err != nil {
		log.Println(err)
	}
	for _, value := range x {
		delete(value, "ID")
		delete(value, "Password")
	}*/
	return c.JSON(userList)
}
