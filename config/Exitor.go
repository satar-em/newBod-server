package config

type AppShutdown interface {
	OnExitApp()
}

var AppShutdowns []AppShutdown

func AddAppShutdowns(shutdown AppShutdown) {
	AppShutdowns = append(AppShutdowns, shutdown)
}

func GetAppShutdowns() *[]AppShutdown {
	return &AppShutdowns
}
