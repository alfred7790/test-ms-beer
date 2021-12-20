package env

import (
	"test-ms-beer/pkg/currencylayer"
	"test-ms-beer/repository"
)

type AppEnv struct {
	Repo            repository.Repository
	CurrencyHandler *currencylayer.Handle
}

func New() *AppEnv {
	return &AppEnv{}
}
