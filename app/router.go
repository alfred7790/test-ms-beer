package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"testproyect/app/api"
	"testproyect/app/env"
	_ "testproyect/docs"
)

func InitRouter(app *env.AppEnv) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	rootApp := gin.Default()

	rootApp.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Up and running...")
	})

	r := rootApp.Group("v1")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api.AddRoutes(r, app)

	return rootApp
}
