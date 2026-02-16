package handler

import (
	"github.com/BreakDown-CS/api-ecommerce/internal/ports"
	"github.com/gofiber/fiber/v2"
)

type StaffHandler struct {
	service ports.StaffServices
}

func NewStaffHandler(service ports.StaffServices) *StaffHandler {
	return &StaffHandler{service: service}
}

func (h *StaffHandler) GetStaffList(c *fiber.Ctx) error {
	result, err := h.service.SearchStaffList()
	if err != nil {
		return err
	}
	return c.JSON(result)
}
