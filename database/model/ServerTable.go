package model

import (
	"errors"
	"gorm.io/gorm/clause"
	"server/database"
)

func init() {
	database.AddTables(&ServerTable{})
}

type ServerTable struct {
	EmamiModel
	Code           string `gorm:"uniqueIndex"`
	NameInDatabase string
}

func (r *ServerTable) TableName() string {
	return "new_bod_server-table"
}
func (r *ServerTable) TableNiceName() string {
	return "Server Builtin Table"
}

func (r *ServerTable) Save() error {
	if r.ID != 0 {
		err := database.GetDB().Save(r).Error
		return err
	}
	err := database.GetDB().Create(r).Error
	return err
}
func (r *ServerTable) Fresh() error {
	if r.ID != 0 {
		return errors.New("there is not Role")
	}
	err := database.GetDB().Preload(clause.Associations).First(r).Error
	return err
}

func (r *ServerTable) SetCreatedByAndSave(CreatedBy *User) error {
	return SetCreatedByAndSave(r, CreatedBy)
}

func (r *ServerTable) SetUpdatedByAndSave(UpdatedBy *User) error {
	return SetUpdatedByAndSave(r, UpdatedBy)
}
