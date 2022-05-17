package auth

import (
	"ssibrahimbas.com/e-trade/http"
	"ssibrahimbas.com/e-trade/internal/user"
)

type AuthModule struct {
	s *Service
	h *Handler
}

func New(us *user.Service, h *http.Client) *AuthModule {
	aService := NewService(us)
	aHandler := NewHandler(aService, h)
	return &AuthModule{
		s: aService,
		h: aHandler,
	}
}

func (a *AuthModule) Init() {
	a.h.InitAllVersions()
}
