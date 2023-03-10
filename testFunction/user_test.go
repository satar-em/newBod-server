package testFunction

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
	"log"
	"server/database"
	"server/database/initDB"
	"server/database/model"
	"testing"
)

func Test(t *testing.T) {
	initDB.InitializeDB()
	getAll()
	//typeTest()
	//encriptTest()
}

func getAll() {
	var userList []model.User
	database.GetDB().Preload(clause.Associations).Find(&userList)

	var roleList []model.RoleUser
	database.GetDB().Preload(clause.Associations).Find(&roleList)

	var detailsList []model.ServerDetails
	database.GetDB().Preload(clause.Associations).Find(&detailsList)

	log.Println("pass")
}
func delRelation() {
	var userList []model.User
	database.GetDB().Preload(clause.Associations).Find(&userList)
	database.GetDB().Model(&userList[0]).Association("Role").Clear()
	log.Println("pass")
}
func typeTest() {
	AdminRole := model.RoleUser{Name: "Administrator", Code: "moderator"}
	AdminRole.ID = 121
	model.SetCreatedByAndSave(&AdminRole, nil)
	AdminSetails := model.ServerDetails{Name: "admin detgails", Summery: "kjasd13 354354"}
	AdminSetails.ID = 4684
	model.SetCreatedByAndSave(&AdminSetails, nil)
	log.Println("pass")
}
func encriptTest() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("emaisdasd@&*(@#!4"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte("emaisdasd@&*(@#!4"))
	println(string(hashedPassword), err == nil)

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte("emaisdasd@&*(@#!4"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte("emaisdasd@&*(@#!4"))
	println(string(hashedPassword), err == nil)

}
