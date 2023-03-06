package database

import "gorm.io/gorm"

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}
func SetDB(db0 *gorm.DB) {
	db = db0
}

var Tables []interface{}

func GetTables() []interface{} {
	return Tables
}
func AddTables(table interface{}) {
	Tables = append(Tables, table)
}
