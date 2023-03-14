package router

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
	"log"
	"server/database"
	"server/database/model"
	"strings"
)

func init() {
	AddApiNeedAuthArray(RoutStruct{Path: "/role", Function: getAllRole, Method: Method_Get})
}

func getAllRole(c *fiber.Ctx) error {
	var roleList []model.RoleUser
	OffsetQuery := c.QueryInt("offset", -1)
	LimitQuery := c.QueryInt("limit", -1)
	if OffsetQuery != -1 && LimitQuery != -1 && LimitQuery != 0 {
		OffsetQuery = OffsetQuery * LimitQuery
	}
	database.GetDB().Preload(clause.Associations).Offset(OffsetQuery).Limit(LimitQuery).Find(&roleList)
	bytes, err := json.Marshal(roleList)
	if err != nil {
		log.Println(err)
	}
	var x []map[string]interface{}
	err = json.Unmarshal(bytes, &x)
	if err != nil {
		log.Println(err)
	}
	exclusionQuery := c.Query("exclusion")
	if exclusionQuery != "" {
		exclusions := strings.Split(exclusionQuery, ",")
		for _, value := range x {
			for _, value2 := range exclusions {
				delete(value, value2)
			}
		}
	}
	return c.JSON(x)
}
