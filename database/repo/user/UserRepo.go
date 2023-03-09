package user

import (
	"server/database"
	"server/database/model"
)

func GetUserByUserAndPass(username string, password string) *model.User {
	var user model.User
	database.GetDB().First(&user, "Username = ?", username)
	if user.ID != 0 {
		if user.CheckPasswordWithBcrypt(password) {
			return &user
		}
	}
	return nil
}
