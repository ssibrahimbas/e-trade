package main

import (
	"ssibrahimbas.com/e-trade/config"
	"ssibrahimbas.com/e-trade/databases"
	"ssibrahimbas.com/e-trade/http"
	"ssibrahimbas.com/e-trade/internal/product"
)

func main() {
	config := config.New()
	databases := databases.New(config)
	http := http.New()
	product := product.New(databases.MongoDB, http)
	product.Init()
	http.Listen("0.0.0.0:3000")
}
