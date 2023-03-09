package model

import (
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
