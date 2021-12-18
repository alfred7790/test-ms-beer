package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testproyect/app/api"
	"testproyect/app/config"
	"testproyect/app/env"
)

func InitRouter(app *env.AppEnv) *gin.Engine {
	if config.Config.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	rootApp := gin.Default()

	rootApp.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Up and running...")
	})

	r := rootApp.Group("v1")
	api.AddRoutesV1(r, app)

	return rootApp
}
