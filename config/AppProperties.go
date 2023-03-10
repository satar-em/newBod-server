package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func init() {
	file, err := os.ReadFile("./application.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(file, &AppProp)
	if err != nil {
		log.Fatal(err)
	}
}

type AppProperties struct {
	DataBase struct {
		DBHost     string `yaml:"dbHost"`
		DBPort     string `yaml:"dbPort"`
		DBName     string `yaml:"dbName"`
		DBUsername string `yaml:"dbUsername"`
		DBPassword string `yaml:"dbPassword"`
		LogFile    struct {
			Name string `yaml:"name"`
			Dest string `yaml:"dest"`
		} `yaml:"logFile"`
	} `yaml:"database"`
	WebServer struct {
		SetupPath string `yaml:"setupPath"`
		Security  struct {
			TokenExpireInMinute int    `yaml:"tokenExpireInMinute"`
			LoginPath           string `yaml:"loginPath"`
			LogoutPath          string `yaml:"logoutPath"`
			UsernameParamName   string `yaml:"usernameParamName"`
			PasswordParamName   string `yaml:"passwordParamName"`
		} `yaml:"security"`
		Port    string `yaml:"port"`
		SSLCrt  string `yaml:"sllCRT"`
		SSLKey  string `yaml:"sslKey"`
		LogFile struct {
			Name string `yaml:"name"`
			Dest string `yaml:"dest"`
		} `yaml:"logFile"`
	} `yaml:"webServer"`
	NeedSetup bool
}

var AppProp AppProperties

func GetAppProperties() *AppProperties {
	return &AppProp
}
