package model

import (
	"gorm.io/gorm"
)

type EmamiModel struct {
	gorm.Model
	CreatedBy       *uint
	CreatedByObject *User `gorm:"foreignKey:CreatedBy" json:"-"`
	UpdatedByObject *User `gorm:"foreignKey:UpdatedBy" json:"-"`
	UpdatedBy       *uint
}
