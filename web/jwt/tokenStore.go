package jwt

import "time"

var authStore []*Authentication

func CheckAuthStore() {
	for index, value := range authStore {
		if value.IsExpire() {
			authStore = append(authStore[:index], authStore[index+1:]...)
			break
		}
	}
}

func GetAuth(id string) *Authentication {
	for index, value := range authStore {
		if value.id == id {
			return authStore[index]
		}
	}
	return nil
}
func DeleteAuth(id string) {
	for index, value := range authStore {
		if value.id == id {
			authStore = append(authStore[:index], authStore[index+1:]...)
			break
		}
	}
}

func GetAuthWithUser(userId uint) *Authentication {
	for index, value := range authStore {
		if value.userId == userId {
			return authStore[index]
		}
	}
	auth := Authentication{userId: userId, id: Conf0.IdGenerator(), expireDur: Conf0.Expire, lastUse: time.Now(), data: map[string]interface{}{}}
	authStore = append(authStore, &auth)
	return &auth
}
