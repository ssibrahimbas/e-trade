package product

import (
	"ssibrahimbas.com/e-trade/databases"
	"ssibrahimbas.com/e-trade/http"
)

type ProductModule struct {
	r *Repository
	s *Service
	h *Handler
}

func New(d *databases.MongoDB, h *http.Client) *ProductModule {
	pRepo := NewRepository(d)
	pService := NewService(pRepo)
	pHandler := NewHandler(pService, h)
	return &ProductModule{
		r: pRepo,
		s: pService,
		h: pHandler,
	}
}

func (p *ProductModule) Init() {
	p.h.InitAllVersions()
}
