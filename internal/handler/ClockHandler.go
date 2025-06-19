package handler

import (
	"github.com/example/internal/model"
	"github.com/example/internal/service"
	"github.com/gofiber/fiber/v2"
)

type ClockHandler struct {
	clockService service.ClockService // Assuming you have a ClockService interface defined
}

func NewClockHandler(clockService service.ClockService) *ClockHandler {
	return &ClockHandler{
		clockService: clockService,
	}
}

func (h *ClockHandler) ClockIn(c *fiber.Ctx) error {
	//put request
	var req model.ClockInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	res, err := h.clockService.ClockIn(c.Context(), c.Locals("user_id").(string), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to clock in",
		})
	}
	return c.JSON(res)
}

func (h *ClockHandler) ClockOut(c *fiber.Ctx) error {
	//put request
	var req model.ClockOutRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	res, err := h.clockService.ClockOut(c.Context(), c.Locals("user_id").(string), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to clock out",
		})
	}
	return c.JSON(res)
}
