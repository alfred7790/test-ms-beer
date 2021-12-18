package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"testproyect/app/config"
	"testproyect/entity"
	"testproyect/pkg/currencylayer"
	"testproyect/pkg/utilities"
	"testproyect/repository"
)

// CreateBeer create a new beer
// @Summary returns details about a new beer was created
// @Description Used to create a new beer
// @Tags Beers
// @Produce json
// @Param data body entity.BeerDTO true "struct to create a new beer"
// @Success 201 {object} entity.Beer
// @Failure 400 {object} utilities.FailureResponse
// @Failure 409 {object} utilities.FailureResponse
// @Failure 500 {object} utilities.FailureResponse
// @Router /v1/beers [POST]
func (app *API) CreateBeer(c *gin.Context) {
	var input *entity.Beer
	err := c.BindJSON(&input)
	if err != nil {
		utilities.Failure("Beer data incorrect", err.Error(), http.StatusBadRequest, c)
		return
	}

	err = app.AppEnv.Repo.SaveBeer(input)
	if err != nil {
		if repository.ValidatePSQLError(err, repository.UniqueViolation) {
			msg := fmt.Sprintf("The sku %s already exists", input.SKU)
			utilities.Failure(msg, err.Error(), http.StatusConflict, c)
			return
		}
		utilities.Failure("Error trying to create a new beer", err.Error(), http.StatusInternalServerError, c)
		return
	}

	c.JSON(http.StatusCreated, input)
}

// UpdateABeer update a beer
// @Summary returns details about a beer was updated
// @Description Used to update the beer's info
// @Tags Beers
// @Produce json
// @Param beerid path int true "Beer ID"
// @Param data body entity.BeerDTO true "struct to update a beer"
// @Success 200 {object} entity.Beer
// @Failure 400 {object} utilities.FailureResponse
// @Failure 404 {object} utilities.FailureResponse
// @Failure 500 {object} utilities.FailureResponse
// @Router /v1/beers/{beerid} [PUT]
func (app *API) UpdateABeer(c *gin.Context) {
	id := c.Param("id")
	beerId, err := strconv.Atoi(id)
	if err != nil {
		utilities.Failure("The BeerID should be a number", err.Error(), http.StatusBadRequest, c)
		return
	}

	var input *entity.Beer
	err = c.BindJSON(&input)
	if err != nil {
		utilities.Failure("Beer data incorrect", err.Error(), http.StatusBadRequest, c)
		return
	}

	input.Id = uint(beerId)

	err = app.AppEnv.Repo.UpdateBeer(input)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utilities.Failure("The beerID doesn't exist", err.Error(), http.StatusNotFound, c)
			return
		}
		utilities.Failure("Error trying to create a new beer", err.Error(), http.StatusInternalServerError, c)
		return
	}

	c.JSON(http.StatusOK, input)
}

// GetBeerByID get a beer info by id
// @Summary returns a beer info searching by beerId
// @Description Used to find a beer
// @Tags Beers
// @Produce json
// @Param beerid path int true "Beer ID"
// @Success 200 {object} entity.Beer
// @Failure 400 {object} utilities.FailureResponse
// @Failure 404 {object} utilities.FailureResponse
// @Failure 500 {object} utilities.FailureResponse
// @Router /v1/beers/{beerid} [GET]
func (app *API) GetBeerByID(c *gin.Context) {
	id := c.Param("id")
	beerId, err := strconv.Atoi(id)
	if err != nil {
		utilities.Failure("The BeerID should be a number", err.Error(), http.StatusBadRequest, c)
		return
	}

	beer, err := app.AppEnv.Repo.GetBeer(uint(beerId))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utilities.Failure("The beerID doesn't exist", err.Error(), http.StatusNotFound, c)
			return
		}
		utilities.Failure("Error trying to get beer's info", err.Error(), http.StatusInternalServerError, c)
		return
	}

	c.JSON(http.StatusOK, beer)
}

// GetBoxPrice get a beer info by id
// @Summary returns a beer info searching by beerId
// @Description List the price of a box of beer
// @Tags Beers
// @Produce json
// @Param beerid path int true "Beer ID"
// @Param currency query string false "Currency, default USD"
// @Param quantity query int false "Quantity, default 6"
// @Success 200 {object} entity.BeerBox
// @Failure 400 {object} utilities.FailureResponse
// @Failure 404 {object} utilities.FailureResponse
// @Failure 500 {object} utilities.FailureResponse
// @Router /v1/beers/{beerid}/boxprice [GET]
func (app *API) GetBoxPrice(c *gin.Context) {
	id := c.Param("id")

	currency := c.DefaultQuery("currency", "USD")
	quantityQry := c.DefaultQuery("quantity", "6")
	quantity, err := strconv.Atoi(quantityQry)
	if err != nil {
		utilities.Failure("The quantity should be a number", err.Error(), http.StatusBadRequest, c)
		return
	}

	beerId, err := strconv.Atoi(id)
	if err != nil {
		utilities.Failure("The BeerID should be a number", err.Error(), http.StatusBadRequest, c)
		return
	}

	beer, err := app.AppEnv.Repo.GetBeer(uint(beerId))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utilities.Failure("The beerID doesn't exist", err.Error(), http.StatusNotFound, c)
			return
		}
	}

	boxPrice := beer.GetBeerBox(uint(quantity))
	boxPrice.AddPrice("USD", utilities.ToFixed(beer.UnitPrice, 2), utilities.ToFixed(float64(quantity)*beer.UnitPrice, 2))

	var resp *currencylayer.CurrencyResponse
	if config.Config.CurrencySubcription {
		resp, err = app.AppEnv.CurrencyHandler.Convert("USD", currency, beer.UnitPrice)
		if err != nil {
			utilities.Failure("currency error", err.Error(), http.StatusBadRequest, c)
			return
		}
		boxPrice.AddPrice(currency, utilities.ToFixed(resp.Info.Quote, 2), utilities.ToFixed(resp.Result, 2))
	} else {
		resp, err = app.AppEnv.CurrencyHandler.Live()
		if err != nil {
			utilities.Failure("currency error", err.Error(), http.StatusBadRequest, c)
			return
		}

		boxPrice.AddPrice("AUD", utilities.ToFixed(resp.Quotes.USDAUD, 2), utilities.ToFixed(beer.UnitPrice*resp.Quotes.USDAUD*float64(quantity), 2))
		boxPrice.AddPrice("CAD", utilities.ToFixed(resp.Quotes.USDCAD, 2), utilities.ToFixed(beer.UnitPrice*resp.Quotes.USDCAD*float64(quantity), 2))
		boxPrice.AddPrice("PLN", utilities.ToFixed(resp.Quotes.USDPLN, 2), utilities.ToFixed(beer.UnitPrice*resp.Quotes.USDPLN*float64(quantity), 2))
		boxPrice.AddPrice("MXN", utilities.ToFixed(resp.Quotes.USDMXN, 2), utilities.ToFixed(beer.UnitPrice*resp.Quotes.USDMXN*float64(quantity), 2))
	}

	c.JSON(http.StatusOK, boxPrice)
}

// ListBeers get a list of beers
// @Summary returns a list of paginated beers
// @Description a page and a limit of results per page is optional.
// @Tags Beers
// @Produce json
// @Param page query string false "Request page, default 1"
// @Param limit query string false "number of results per page, default 50"
// @Success 200 {object} entity.BeerList
// @Failure 500 {object} utilities.FailureResponse
// @Router /v1/beers [GET]
func (app *API) ListBeers(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))
	if err != nil || limit <= 0 {
		limit = 50
	}

	if limit > 100 {
		limit = 100
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page == 0 {
		page = 1
	}

	total, beers, err := app.AppEnv.Repo.GetBeers(page, limit)
	if err != nil {
		utilities.Failure("Error trying to get a list of beers", err.Error(), http.StatusInternalServerError, c)
		return
	}

	results := entity.BeerList{
		Total: total,
		Beers: beers,
	}
	c.JSON(http.StatusOK, results)
}
