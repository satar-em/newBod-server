package model

import (
	"gorm.io/gorm"
)

type EmamiModel struct {
	gorm.Model
	CreatedBy uint
	UpdatedBy uint
}

func (m EmamiModel) GetCreatedBy(db *gorm.DB) User {
	user := User{}
	db.First(&user, m.CreatedBy)
	return user
}
func (m EmamiModel) GetUpdatedBy(db *gorm.DB) User {
	user := User{}
	db.First(&user, m.UpdatedBy)
	return user
}
