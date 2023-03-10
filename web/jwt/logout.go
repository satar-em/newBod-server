package jwt

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetLogout(c *fiber.Ctx) error {
	if c.GetReqHeaders()["Authentication"] != "" {
		auth := GetAuth(c.GetReqHeaders()["Authentication"])
		DeleteAuth(auth.id)
		return c.JSON(LogoutResponse{Message: fmt.Sprintf("success logout user(%d)", *auth.GetUser()), success: true})
	}
	return c.JSON(LogoutResponse{Message: "failed to logout ", success: false})
}

type LogoutResponse struct {
	Message string `json:"message"`
	success bool
}
