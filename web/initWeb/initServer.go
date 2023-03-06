package initWeb

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"server/config"
)

func InitWebserver() {
	WebApp := fiber.New()
	WebApp.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "content-type,authentication",
		AllowCredentials: true,
		ExposeHeaders:    "*",
		MaxAge:           0,
	}))
	createLogger(WebApp)
	initStartForFistTime(WebApp)
	WebApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	log.Fatal(WebApp.ListenTLS(":"+config.GetAppProperties().WebServer.Port, config.GetAppProperties().WebServer.SSLCrt, config.GetAppProperties().WebServer.SSLKey))
}

func createLogger(app *fiber.App) {
	logFileSource := fmt.Sprintf("%s/%s",
		config.GetAppProperties().WebServer.LogFile.Dest,
		config.GetAppProperties().WebServer.LogFile.Name)
	file, err := os.OpenFile(logFileSource, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	app.Use(logger.New(logger.Config{
		Output: file,
		Format: "[${time}] (${ip}) ${status} - ${latency} ${method} ${path}\n",
	}))
	config.AddAppShutdowns(ShutDownLogFile{file})
}

type ShutDownLogFile struct {
	LogFile *os.File
}

func (s ShutDownLogFile) OnExitApp() {
	println("closing LogFile of fiber")
	err := s.LogFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
