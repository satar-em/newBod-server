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
	AddApiNeedAuthArray(RoutStruct{Path: "/serverTable", Function: getAllTables, Method: Method_Get})
}

func getAllTables(c *fiber.Ctx) error {
	var serverTablesList []model.ServerTable
	OffsetQuery := c.QueryInt("offset", -1)
	LimitQuery := c.QueryInt("limit", -1)
	if OffsetQuery != -1 && LimitQuery != -1 && LimitQuery != 0 {
		OffsetQuery = OffsetQuery * LimitQuery
	}
	database.GetDB().Preload(clause.Associations).Offset(OffsetQuery).Limit(LimitQuery).Find(&serverTablesList)
	bytes, err := json.Marshal(serverTablesList)
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
