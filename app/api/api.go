package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testproyect/app/env"
)

type API struct {
	AppEnv *env.AppEnv
}

func AddRoutes(r *gin.RouterGroup, app *env.AppEnv)  {
	api := API{
		AppEnv: app,
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "V1 is running")
	})

	beers := r.Group("beers")
	beers.GET("", api.ListBeers)
	beers.POST("", api.CreateBeer)
	beers.GET(":id", api.GetBeerByID)
	beers.PUT(":id", api.UpdateABeer)
	beers.GET(":id/boxprice", api.GetBoxPrice)

}