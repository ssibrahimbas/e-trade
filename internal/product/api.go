package product

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ssibrahimbas.com/e-trade/tools"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	p := &CreateProductDto{}
	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(tools.Error("Something went wrong"))
	}
	res, err := h.s.Create(p)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(tools.Error(err.Error()))
	}
	return c.Status(fiber.StatusCreated).JSON(tools.SuccessData("Product successfully created.", res))
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	res, err := h.s.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(tools.Error(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(tools.SuccessData("Products successfully fetched.", res))
}

func (h *Handler) GetDetail(c *fiber.Ctx) error {
	params := c.Params("id")
	_id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(tools.Error("Something went wrong"))
	}
	res, err := h.s.GetDetail(&GetDetailProductDto{Id: _id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(tools.Error(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(tools.SuccessData("Product successfully fetched.", res))
}
