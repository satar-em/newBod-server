package initWeb

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/utils"
	"log"
	"os"
	"server/config"
	"server/web/jwt"
	"server/web/router"
	"time"
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
	WebApp.Use(jwt.Set(jwt.AuthConfig{
		IdGenerator: utils.UUIDv4,
		LoginPath:   config.GetAppProperties().WebServer.Security.LoginPath,
		LogoutPath:  config.GetAppProperties().WebServer.Security.LogoutPath,
		Expire:      time.Duration(config.GetAppProperties().WebServer.Security.TokenExpireInMinute) * time.Minute,
	}))
	initRouterArray(WebApp)
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

func initRouterArray(app *fiber.App) {

	for _, value := range router.GetRoutArray() {
		switch value.Metod {
		case router.Method_Get:
			app.Get(value.Path, value.Function)
			return
		case router.Method_Post:
			app.Post(value.Path, value.Function)
			return
		case router.Method_Put:
			app.Put(value.Path, value.Function)
			return
		case router.Method_Delete:
			app.Delete(value.Path, value.Function)
			return
		}
	}
}
