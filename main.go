package main

import (
	"fmt"
	appCtx "test-ms-beer/app"
	"test-ms-beer/app/config"
	"test-ms-beer/app/env"
	"test-ms-beer/pkg/currencylayer"
	"test-ms-beer/pkg/utilities"
	"test-ms-beer/repository"
)

// @title Test API
// @version 1.0
// @description test api
// @BasePath /
func main() {
	app := env.New()
	app.Repo = repository.NewPostgresRepository()
	app.CurrencyHandler = currencylayer.NewHandle()
	app.CurrencyHandler.SetHostAndKey(
		config.Config.CurrencyHost,
		config.Config.CurrencyAccessKey)
	utilities.SetAllowInsecureCert(config.Config.AllowInsecureCert)

	go initDB(app.Repo)

	r := appCtx.InitRouter(app)
	err := r.Run(fmt.Sprintf(":%s", config.Config.Port))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func initDB(repo repository.Repository) error {
	c := config.Config
	if err := repo.Init(c.DBIP, c.DBPort, c.DBUser, c.DBPass, c.DBName, c.DBRetryCount); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
