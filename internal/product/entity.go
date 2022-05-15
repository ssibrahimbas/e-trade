package product

import "time"

type Product struct {
	Id         string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	StockCount int       `json:"stockCount"`
	Date       time.Time `json:"date"`
}
