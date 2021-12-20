package api

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"test-ms-beer/app/env"
	"test-ms-beer/entity"
	"test-ms-beer/pkg/models"
	"test-ms-beer/pkg/utilities"
	"testing"
)

func TestAPI_ListBeers(t *testing.T) {
	beersMock := []*entity.Beer{
		&entity.Beer{
			Id:        1,
			SKU:       "BEERSKU1",
			Name:      "Beer1",
			UnitPrice: 1.2,
		},
		&entity.Beer{
			Id:        2,
			SKU:       "BEERSKU2",
			Name:      "Beer2",
			UnitPrice: 1.19,
		},
		&entity.Beer{
			Id:        3,
			SKU:       "BEERSKU3",
			Name:      "Beer3",
			UnitPrice: 1.90,
		},
		&entity.Beer{
			Id:        4,
			SKU:       "BEERSKU4",
			Name:      "Beer3",
			UnitPrice: 0.89,
		},
	}

	tests := []struct {
		expectedBody interface{}
		expectedCode int
		mockFunc     func(page, limit int) (uint, []*entity.Beer, error)
		expectError  bool
	}{
		{
			entity.BeerList{
				Total: 0,
				Beers: []*entity.Beer{},
			},
			http.StatusOK,
			func(page, limit int) (uint, []*entity.Beer, error) {
				return 0, []*entity.Beer{}, nil
			},
			false,
		},
		{
			entity.BeerList{
				Total: 2,
				Beers: beersMock[:2],
			},
			http.StatusOK,
			func(page, limit int) (uint, []*entity.Beer, error) {
				return 2, beersMock[:2], nil
			},
			false,
		},
		{
			utilities.FailureResponse{
				Message: "Error trying to get a list of beers",
				Details: "internal error mock",
				Status:  "failure",
			},
			http.StatusInternalServerError,
			func(page, limit int) (uint, []*entity.Beer, error) {
				return 0, []*entity.Beer{}, errors.New("internal error mock")
			},
			true,
		},
	}

	for _, test := range tests {
		appEnv := env.New()
		mockRepo := &models.RepositoryMock{}
		appEnv.Repo = mockRepo
		mockRepo.GetBeersMockFn = test.mockFunc
		api := API{
			AppEnv: appEnv,
		}
		// gin.SetMode(gin.DebugMode)
		rootApp := gin.Default()
		r := rootApp.Group("v1")
		beers := r.Group("beers")
		beers.GET("", api.ListBeers)
		req, _ := http.NewRequest(http.MethodGet, "/v1/beers", nil)
		recorder := httptest.NewRecorder()
		rootApp.ServeHTTP(recorder, req)

		if test.expectedCode != recorder.Code {
			t.Fatalf("unexpected HTTP response code.  Expected %d, got: %d", test.expectedCode, recorder.Code)
		}

		if test.expectError {
			var failureResponse utilities.FailureResponse
			err := json.Unmarshal(recorder.Body.Bytes(), &failureResponse)
			if err != nil {
				t.Errorf("error unmarshalling response data: %v", err.Error())
			}

			assert.Equal(t, test.expectedBody, failureResponse)
		} else {
			var successResponse entity.BeerList
			err := json.Unmarshal(recorder.Body.Bytes(), &successResponse)
			if err != nil {
				t.Errorf("error unmarshalling response data: %v", err.Error())
			}

			assert.Equal(t, test.expectedBody, successResponse)
		}
	}
}
