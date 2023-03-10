package jwt

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type AuthConfig struct {
	IdGenerator func() string
	Expire      time.Duration
	LoginPath   string
	LogoutPath  string
}

var Conf0 AuthConfig

func Set(conf AuthConfig) func(c *fiber.Ctx) error {
	Conf0 = conf
	emami := func(c *fiber.Ctx) error {
		CheckAuthStore()
		if c.Path() == Conf0.LoginPath && c.Method() == "POST" {
			return PostLogin(c)
		}
		if c.GetReqHeaders()["Authentication"] != "" {
			auth := GetAuth(c.GetReqHeaders()["Authentication"])
			if auth != nil {
				auth.Fresh()
				if c.Path() == Conf0.LogoutPath && c.Method() == "GET" {
					return GetLogout(c)
				}
				return c.Next()
			}
		}
		c.Status(401)
		return c.JSON(NotAuthentication{Code: 401, Message: "authentication fail :("})
	}
	return emami
}

type NotAuthentication struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
