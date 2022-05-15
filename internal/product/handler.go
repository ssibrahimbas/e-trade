package product

import (
	"ssibrahimbas.com/e-trade/http"
)

type Handler struct {
	s *Service
	h *http.Client
}

func NewHandler(s *Service, h *http.Client) *Handler {
	return &Handler{
		s: s,
		h: h,
	}
}

func (h *Handler) InitAllVersions() {
	h.initV1()
}

func (h *Handler) initV1() {
	v1 := h.h.App.Group("v1/products")
	v1.Get("/", h.GetAll)
	v1.Get("/:id", h.GetDetail)
	v1.Post("/create", h.Create)
}
