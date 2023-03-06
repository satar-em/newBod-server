package model

import (
	"github.com/lib/pq"
	"server/database"
)

func init() {
	database.AddTables(&User{})
}

type User struct {
	EmamiModel
	Role        []*RoleUser `gorm:"many2many:new_bod_user_role;"`
	Name        string
	Email       string
	PhoneNumber string
	Address     string
	UserName    string
	Password    string
	ErpEmami    pq.Int64Array `gorm:"type:integer[]"`
}

func (u User) TableName() string {
	return "new_bod_user"
}
