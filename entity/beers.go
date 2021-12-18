package entity

type Beer struct {
	Id        uint    `json:"id" gorm:"primary_key" example:"1" descripcion:"Primary key of the beer"`
	SKU       string  `json:"sku" gorm:"unique" example:"BHE001" description:"Unique key for the beer"`
	Name      string  `json:"name" gorm:"not null" example:"Bohemia" description:"Name of the beer"`
	UnitPrice float64 `json:"unitPrice" gorm:"not null" example:"1.20" description:"Price of the beer of curse, in USD"`
}

type BeerBox struct {
	Beer `json:"beer"`

	Quantity uint            `json:"quantity" example:"10"`
	Prices   []*BeerBoxPrice `json:"prices"`
}

type BeerBoxPrice struct {
	CurrencyName  string  `json:"currencyName" example:"USD"`
	CurrencyPrice float64 `json:"currencyPrice" example:"1.20"`
	Total         float64 `json:"total" example:"12"`
}

type BeerDTO struct {
	SKU       string  `json:"sku" example:"BHE001" description:"Unique key for the beer"`
	Name      string  `json:"name" example:"Bohemia" description:"Name of the beer"`
	UnitPrice float64 `json:"unitPrice" example:"1.20" description:"Price of the beer of curse in USD"`
}

type BeerList struct {
	Total uint    `json:"total" example:"100"`
	Beers []*Beer `json:"beers"`
}

func (b *Beer) GetBeerBox(quantity uint) *BeerBox {
	return &BeerBox{
		Beer:     *b,
		Quantity: quantity,
		Prices:   make([]*BeerBoxPrice, 0),
	}
}

func (b *BeerBox) AddPrice(currencyName string, price, total float64) {
	b.Prices = append(b.Prices, &BeerBoxPrice{
		CurrencyName: currencyName,
		CurrencyPrice: price,
		Total: total,
	})
}
