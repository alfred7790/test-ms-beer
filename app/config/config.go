package config

import (
	"fmt"
)
import "github.com/jinzhu/configor"

// Settings Some of this field should be loaded from a environment file (.env)
type Settings struct {
	// APP environments
	AppName string `default:"test"`
	Port    string `default:"8080"`

	// DB environments
	DBName       string `default:"test"`
	DBIP         string `default:"localhost"`
	DBPort       string `default:"5432"`
	DBUser       string `default:"tester"`
	DBPass       string `default:"mySuperPass"`
	DBRetryCount int    `default:"1"`

	// CurrencyLayer API environments
	CurrencyHost        string `default:"http://api.currencylayer.com"`
	CurrencyAccessKey   string `default:"76d037eb3c53eaeb6b9e6ddc54463bd3"`
	CurrencySubcription bool   `default:"false"`
	AllowInsecureCert   bool   `default:"true"`
}

var Config = Settings{}

func init() {
	if err := configor.Load(&Config, "config.yml"); err != nil {
		fmt.Println("Error trying to load configuration", err.Error())
	}
}
