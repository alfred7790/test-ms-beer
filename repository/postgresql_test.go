package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
	"testproyect/entity"
)

func TestRepo_GetBeer(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("failed to open sqlmock database: %v", err)
	}

	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		t.Errorf("failed to open gorm mock database: %v", err)
	}
	defer db.Close()

	expected := &entity.Beer{
		Id:        1,
		SKU:       "BEERSKU",
		Name:      "Beer",
		UnitPrice: 1.2,
	}

	rows := sqlmock.NewRows([]string{"id", "sku", "name", "unit_price"}).
		AddRow(expected.Id, expected.SKU, expected.Name, expected.UnitPrice)

	mock.ExpectQuery(`SELECT * FROM "beers" WHERE ("beers"."id" = $1) ORDER BY "beers"."id" ASC LIMIT 1`).
		WithArgs(expected.Id).
		WillReturnRows(rows)

	sut := NewPostgresRepository()
	sut.SQLDB = gormDB
	sut.DBName = "test"
	actual, err := sut.GetBeer(expected.Id)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, actual)
}

func TestRepo_GetBeers(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("failed to open sqlmock database: %v", err)
	}

	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		t.Errorf("failed to open gorm mock database: %v", err)
	}
	defer db.Close()

	page := 1
	limit := 50

	expectedBeers := []*entity.Beer{
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

	rows := sqlmock.NewRows([]string{"id", "sku", "name", "unit_price"})

	for _, row := range expectedBeers {
		rows.AddRow(row.Id, row.SKU, row.Name, row.UnitPrice)

	}

	mock.ExpectQuery(`SELECT * FROM "beers" LIMIT 50 OFFSET 0`).
		WillReturnRows(rows)

	sut := NewPostgresRepository()
	sut.SQLDB = gormDB
	sut.DBName = "test"
	_, beers, err := sut.GetBeers(page, limit)
	if err != nil {
		t.Error(err)
	}

	for i, v := range beers {
		assert.Equal(t, expectedBeers[i], v)
	}
}
