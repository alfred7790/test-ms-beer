package models

import (
	"testproyect/entity"
	"testproyect/repository"
)

type RepositoryMock struct {
	repository.BeersRepo

	InitMockFn       func(dbIP, dbPort, dbUser, dbPass, dbName string, retryCount int) error
	SaveBeerMockFn   func(beer *entity.Beer) error
	GetBeersMockFn   func(page, limit int) (uint, []*entity.Beer, error)
	GetBeerMockFn    func(beerId uint) (*entity.Beer, error)
	UpdateBeerMockFn func(beer *entity.Beer) error
	GetBeersCalled   bool
	SaveBeerCalled   bool
	GetBeerCalled    bool
	UpdateBeerCalled bool
}

func NewRepoMock() *RepositoryMock {
	return &RepositoryMock{}
}

func (m *RepositoryMock) Init(dbIP, dbPort, dbUser, dbPass, dbName string, retryCount int) error {
	return m.InitMockFn(dbIP, dbPort, dbUser, dbPass, dbName, retryCount)
}

func (m *RepositoryMock) SaveBeer(beer *entity.Beer) error {
	m.SaveBeerCalled = true
	return m.SaveBeerMockFn(beer)
}

func (m *RepositoryMock) GetBeers(page, limit int) (uint, []*entity.Beer, error) {
	m.GetBeersCalled = true
	return m.GetBeersMockFn(page, limit)
}

func (m *RepositoryMock) GetBeer(beerId uint) (*entity.Beer, error) {
	m.GetBeerCalled = true
	return m.GetBeerMockFn(beerId)
}

func (m *RepositoryMock) UpdateBeer(beer *entity.Beer) error {
	m.UpdateBeerCalled = true
	return m.UpdateBeerMockFn(beer)
}
