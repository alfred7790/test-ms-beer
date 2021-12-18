package env

import (
	"testproyect/pkg/currencylayer"
	"testproyect/repository"
)

type AppEnv struct {
	Repo            repository.Repository
	CurrencyHandler *currencylayer.Handle
}

func New() *AppEnv {
	return &AppEnv{}
}
