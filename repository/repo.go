package repository

import "test-ms-beer/entity"

type Repo struct {
	Base
}

type Repository interface {
	Init(dbIP, dbPort, dbUser, dbPass, dbName string, retryCount int) error
	BeersRepo
}

type BeersRepo interface {
	SaveBeer(beer *entity.Beer) error
	GetBeers(page, limit int) (uint, []*entity.Beer, error)
	GetBeer(beerId uint) (*entity.Beer, error)
	UpdateBeer(beer *entity.Beer) error
}
