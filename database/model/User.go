package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"server/database"
)

func init() {
	database.AddTables(&User{})
}

type User struct {
	gorm.Model
	CreatedBy0       *uint       `json:"CreatedBy"`
	CreatedByObject0 *User       `gorm:"foreignKey:CreatedBy0" json:"-"`
	UpdatedByObject0 *User       `gorm:"foreignKey:UpdatedBy0" json:"-"`
	UpdatedBy0       *uint       `json:"UpdatedBy"`
	Role             []*RoleUser `gorm:"many2many:new_bod_user-role;"`
	Password         string      `gorm:"-"`
	PasswordEncrypt  string      `json:"-"`
	Username         string      `gorm:"uniqueIndex"`
	Name             string
	Email            string
	PhoneNumber      string
	Address          string
}

func (u *User) TableName() string {
	return "new_bod_user"
}

func (u *User) SetPasswordWithBcrypt(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err == nil {
		u.PasswordEncrypt = string(hashedPassword)
	}
	return err
}
func (u *User) CheckPasswordWithBcrypt(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordEncrypt), []byte(password))
	return err == nil
}

func (u *User) Save() error {
	u.Password = ""
	if u.ID != 0 {
		err := database.GetDB().Save(u).Error
		return err
	}
	return database.GetDB().Create(u).Error
}

func (u *User) Fresh() error {
	if u.ID == 0 {
		return errors.New("there is not User")
	}
	return database.GetDB().Preload(clause.Associations).First(u).Error
}
func (u *User) SetRole(roles []*RoleUser) error {
	if u.ID == 0 {
		return errors.New("there is not User")
	}
	return database.GetDB().Model(u).Association("Role").Replace(roles)
}
func (u *User) SetCreatedByAndSave(CreatedBy *User) error {
	if u.ID == 0 {
		return errors.New("there is not User")
	}
	return database.GetDB().Model(u).Association("CreatedByObject0").Replace(CreatedBy)
}
func (u *User) SetUpdatedByAndSave(UpdatedBy *User) error {
	if u.ID == 0 {
		return errors.New("there is not User")
	}
	return database.GetDB().Model(u).Association("UpdatedByObject0").Replace(UpdatedBy)
}
