package auth

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
	v1 := h.h.App.Group("v1/auth")
	v1.Post("/login", h.Login)
	v1.Post("/register", h.Register)
	v1.Post("/forgot/password", h.ForgotPassword)
	v1.Post("/forgot/username", h.ForgotUsername)
}
