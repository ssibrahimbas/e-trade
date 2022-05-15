package http

import "github.com/gofiber/fiber/v2"

type Client struct {
	App *fiber.App
}

func New() *Client {
	return &Client{
		App: fiber.New(),
	}
}

func (h *Client) Listen(p string) error {
	return h.App.Listen(p)
}
