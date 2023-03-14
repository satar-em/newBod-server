package model

import (
	"errors"
	"gorm.io/gorm/clause"
	"server/database"
)

func init() {
	database.AddTables(&RoleUser{})
}

type RoleUser struct {
	EmamiModel
	Name        string
	Code        string  `gorm:"uniqueIndex"`
	UserContain []*User `gorm:"many2many:new_bod_user-role;"`
}

func (r *RoleUser) TableName() string {
	return "new_bod_role"
}
func (r *RoleUser) TableNiceName() string {
	return "Role Fro User"
}

func (r *RoleUser) Save() error {
	if r.ID != 0 {
		err := database.GetDB().Save(r).Error
		return err
	}
	err := database.GetDB().Create(r).Error
	return err
}
func (r *RoleUser) Fresh() error {
	if r.ID != 0 {
		return errors.New("there is not Role")
	}
	err := database.GetDB().Preload(clause.Associations).First(r).Error
	return err
}

func (r *RoleUser) SetCreatedByAndSave(CreatedBy *User) error {
	return SetCreatedByAndSave(r, CreatedBy)
}

func (r *RoleUser) SetUpdatedByAndSave(UpdatedBy *User) error {
	return SetUpdatedByAndSave(r, UpdatedBy)
}
