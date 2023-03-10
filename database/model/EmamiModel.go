package model

import (
	"errors"
	"gorm.io/gorm"
	"reflect"
	"server/database"
)

type EmamiModel struct {
	gorm.Model
	CreatedBy       *uint
	CreatedByObject *User `gorm:"foreignKey:CreatedBy" json:"-"`
	UpdatedByObject *User `gorm:"foreignKey:UpdatedBy" json:"-"`
	UpdatedBy       *uint
}

func SetCreatedByAndSave(model interface{}, CreatedBy *User) error {
	ID := reflect.ValueOf(model).Elem().FieldByName("ID").Uint()
	if ID == 0 {
		return errors.New("there is not User")
	}
	err := database.GetDB().Model(model).Association("CreatedByObject").Replace(CreatedBy)
	return err
}

func SetUpdatedByAndSave(model interface{}, UpdatedBy *User) error {
	ID := reflect.ValueOf(model).Elem().FieldByName("ID").Uint()
	if ID == 0 {
		return errors.New("there is not User")
	}
	err := database.GetDB().Model(model).Association("UpdatedByObject").Replace(UpdatedBy)
	return err
}

type EmamiModelFunctions interface {
	SetCreatedByAndSave(CreatedBy *User) error
	SetUpdatedByAndSave(UpdatedBy *User) error
}
