package product

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateProductDto struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	StockCount int     `json:"stockCount"`
}

type GetDetailProductDto struct {
	Id primitive.ObjectID `json:"id"`
}
