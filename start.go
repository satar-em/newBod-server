package main

import (
	"fmt"
	"os"
	"os/signal"
	"server/config"
	"server/database/initDB"
	"server/web/initWeb"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			for index, value := range *config.GetAppShutdowns() {
				fmt.Printf("\n\n*****  Shutting Down %d of %d\n", index+1, len(*config.GetAppShutdowns()))
				value.OnExitApp()
			}
			os.Exit(1)
		}
	}()
	initDB.InitializeDB()
	initWeb.InitWebserver()
}
