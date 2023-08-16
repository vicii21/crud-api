package models

import (
	"time"
)

type Product struct {
	ProductID  int64     `json:"productid"`
	Name       string    `json:"name"`
	ShortDesc  string    `json:"shortdesc"`
	Desc       string    `json:"desc"`
	Price      float64   `json:"price"`
	Quantity   int64     `json:"quantity"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
	CategoryID int64     `json:"categoryid"`
}
