package main

import (
	"ssibrahimbas.com/e-trade/config"
	"ssibrahimbas.com/e-trade/databases"
	"ssibrahimbas.com/e-trade/http"
	"ssibrahimbas.com/e-trade/internal/auth"
	"ssibrahimbas.com/e-trade/internal/product"
	"ssibrahimbas.com/e-trade/internal/user"
	"ssibrahimbas.com/e-trade/services"
)

func main() {
	config := config.New()
	databases := databases.New(config)
	http := http.New()
	cipher := services.NewCipher()

	user := user.New(cipher, databases.MongoDB)
	auth.New(user.S, http).Init()
	product.New(databases.MongoDB, http).Init()
	http.Listen("0.0.0.0:3000")
}
