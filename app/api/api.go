package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"testproyect/app/env"
	_ "testproyect/docs"
)

type API struct {
	AppEnv *env.AppEnv
}

func AddRoutesV1(r *gin.RouterGroup, app *env.AppEnv) {
	api := API{
		AppEnv: app,
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "V1 is running")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	beers := r.Group("beers")
	beers.GET("", api.ListBeers)
	beers.POST("", api.CreateBeer)
	beers.GET(":id", api.GetBeerByID)
	beers.PUT(":id", api.UpdateABeer)
	beers.GET(":id/boxprice", api.GetBoxPrice)

}
