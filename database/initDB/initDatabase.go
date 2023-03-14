package initDB

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"server/config"
	"server/database"
	"server/database/model"
	"time"
)

func InitializeDB() {
	strConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.GetAppProperties().DataBase.DBHost,
		config.GetAppProperties().DataBase.DBUsername,
		config.GetAppProperties().DataBase.DBPassword,
		config.GetAppProperties().DataBase.DBName,
		config.GetAppProperties().DataBase.DBPort,
	)

	db, err := gorm.Open(postgres.Open(strConnection), &gorm.Config{Logger: createLoggerFile()})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(database.GetTables()...)
	if err != nil {
		log.Fatal(err)
	}

	for _, value0 := range database.GetTables() {
		value, ok := value0.(database.EmamiTabler)
		if !ok {
			log.Fatal("err")
		}
		table := model.ServerTable{Code: value.TableNiceName(), NameInDatabase: value.TableName()}
		err0 := db.Create(&table).Error
		if err0 != nil {
			log.Println(err0)
		}
	}
	var user1 []model.User
	db.Preload("Role.UserContain.Role.UserContain").Preload(clause.Associations).Find(&user1)
	if len(user1) == 0 {
		config.GetAppProperties().NeedSetup = true
	} else {
		config.GetAppProperties().NeedSetup = false
	}
	database.SetDB(db)

}

func TestMain() {
	strConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.GetAppProperties().DataBase.DBHost,
		config.GetAppProperties().DataBase.DBUsername,
		config.GetAppProperties().DataBase.DBPassword,
		config.GetAppProperties().DataBase.DBName,
		config.GetAppProperties().DataBase.DBPort,
	)
	db, err := gorm.Open(postgres.Open(strConnection), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to database ", db.Name())
	err = db.AutoMigrate(&model.User{})
	err = db.AutoMigrate(&model.RoleUser{})
	user1 := model.User{}
	//role1 := model.RoleUser{}
	db.Preload("Role.UserContain.Role.UserContain").Preload(clause.Associations).First(&user1, 1)

	//user1.Role[0].CreatedBy = 0
	//db.Save(&user1.Role[0])

	/*db.Preload("UserContain").First(&role1, 4)
	erpss, _ := json.Marshal(role1)
	println(string(erpss))*/

	/*var user1 []model.User
	db.Preload("Role.UserContain.Role.UserContain").Preload(clause.Associations).Find(&user1)
	fmt.Printf("size of found user is %d .\n", len(user1))*/
}

func createLoggerFile() logger.Interface {
	logFileSource := fmt.Sprintf("%s/%s",
		config.GetAppProperties().DataBase.LogFile.Dest,
		config.GetAppProperties().DataBase.LogFile.Name)
	LogFile, _ := os.OpenFile(logFileSource, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	newLogger := logger.New(
		log.New(LogFile, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	erp := ShutDownLogFile{LogFile}
	config.AddAppShutdowns(erp)
	return newLogger
}

type ShutDownLogFile struct {
	LogFile *os.File
}

func (s ShutDownLogFile) OnExitApp() {
	println("closing LogFile of Gorm")
	err := s.LogFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
