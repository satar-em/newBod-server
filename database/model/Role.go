package model

import "server/database"

func init() {
	database.AddTables(&RoleUser{})
}

type RoleUser struct {
	EmamiModel
	Name        string
	Code        string
	UserContain []*User `gorm:"many2many:new_bod_user-role;"`
}

func (r *RoleUser) TableName() string {
	return "new_bod_role"
}
