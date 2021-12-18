package config

import (
	"fmt"
)
import "github.com/jinzhu/configor"

type Settings struct {
	// APP environments
	AppName string `default:"test"`
	Port    string `default:"8080"`

	// DB environments
	DBName       string `default:"test"`
	DBIP         string `default:"localhost"`
	DBPort       string `default:"5432"`
	DBUser       string
	DBPass       string
	DBRetryCount int `default:"-1"`

	// CurrencyLayer API environments
	CurrencyHost        string
	CurrencyAccessKey   string
	CurrencySubcription bool `default:"false"`
	AllowInsecureCert   bool
}

var Config = Settings{}

func init() {
	if err := configor.Load(&Config, "config.yml"); err != nil {
		fmt.Println("Error trying to load configuration", err.Error())
	}
}
