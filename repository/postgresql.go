package repository

import (
	"test-ms-beer/entity"
)

func (r *Repo) SaveBeer(beer *entity.Beer) error {
	err := r.SQLDB.Create(&beer).Scan(&beer).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateBeer(beer *entity.Beer) error {
	err := r.SQLDB.Model(&entity.Beer{Id: beer.Id}).
		Update(&entity.Beer{
			SKU:       beer.SKU,
			Name:      beer.Name,
			UnitPrice: beer.UnitPrice,
		}).Scan(beer).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetBeers(page, limit int) (uint, []*entity.Beer, error) {
	beers := make([]*entity.Beer, 0)
	total := 0
	start := limit * (page - 1)
	r.SQLDB.Model(&entity.Beer{}).Count(&total)
	err := r.SQLDB.Offset(start).Limit(limit).Find(&beers).Error
	return uint(total), beers, err

}
func (r *Repo) GetBeer(beerId uint) (*entity.Beer, error) {
	var beerFound entity.Beer
	err := r.SQLDB.Where(&entity.Beer{Id: beerId}).First(&beerFound).Error
	return &beerFound, err
}
