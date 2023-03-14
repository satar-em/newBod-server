package testFunction

import (
	"gorm.io/gorm/clause"
	"log"
	"server/database"
	"server/database/model"
)

func getAllTable() {
	var serverTablesList []model.ServerTable
	database.GetDB().Preload(clause.Associations).Find(&serverTablesList)

	log.Println("pass")
}
