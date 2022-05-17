package auth

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"ssibrahimbas.com/e-trade/internal/user"
	"ssibrahimbas.com/e-trade/services"
	"ssibrahimbas.com/e-trade/tools"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	u := &user.LoginUserDto{}
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(tools.Error("Something went wrong"))
	}
	errors := services.ValidateStruct(*u)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(tools.ErrorData("Validation failed.", errors))
	}
	res, err := h.s.Login(u)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(tools.Error("User not found."))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(tools.Error(err.Error()))
	}
	return c.Status(fiber.StatusCreated).JSON(tools.SuccessData("Login successfully.", res))
}

func (h *Handler) Register(c *fiber.Ctx) error {
	u := &user.CreateUserDto{}
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(tools.Error("Something went wrong"))
	}
	errors := services.ValidateStruct(*u)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(tools.ErrorData("Validation failed.", errors))
	}
	res, err := h.s.Register(u, c.IP())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(tools.Error(err.Error()))
	}
	return c.Status(fiber.StatusCreated).JSON(tools.SuccessData("Register successfully.", res))
}

func (h *Handler) ForgotUsername(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(tools.Success("Hello world!"))
}

func (h *Handler) ForgotPassword(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(tools.Success("Hello world!"))
}
