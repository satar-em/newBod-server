package model

import (
	"errors"
	"gorm.io/gorm/clause"
	"server/database"
)

func init() {
	database.AddTables(&ServerDetails{})
}

type ServerDetails struct {
	EmamiModel
	Name    string
	Summery string
	Details string
}

func (s *ServerDetails) TableName() string {
	return "new_bod_server-details"
}
func (s *ServerDetails) Save() error {
	if s.ID != 0 {
		err := database.GetDB().Save(s).Error
		return err
	}
	err := database.GetDB().Create(s).Error
	return err
}

func (s *ServerDetails) Fresh() error {
	if s.ID != 0 {
		return errors.New("there is not ServerDetails")
	}
	err := database.GetDB().Preload(clause.Associations).First(s).Error
	return err
}

func (s *ServerDetails) SetCreatedByAndSave(CreatedBy *User) error {
	return SetCreatedByAndSave(s, CreatedBy)
}

func (s *ServerDetails) SetUpdatedByAndSave(UpdatedBy *User) error {
	return SetUpdatedByAndSave(s, UpdatedBy)
}
