package user

import (
	"ssibrahimbas.com/e-trade/databases"
	"ssibrahimbas.com/e-trade/services"
)

type UserModule struct {
	r *Repository
	S *Service
}

func New(c *services.Cipher, d *databases.MongoDB) *UserModule {
	uRepo := NewRepository(d, c)
	uService := NewService(uRepo)
	return &UserModule{
		r: uRepo,
		S: uService,
	}
}
