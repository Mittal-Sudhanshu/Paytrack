package handler

import (
	"github.com/example/internal/model"
	"github.com/example/internal/service"
	"github.com/gofiber/fiber/v2"
)

type InviteHandler struct {
	inviteService service.InviteService
}

func NewInviteHandler(inviteService service.InviteService) *InviteHandler {
	return &InviteHandler{inviteService: inviteService}
}

func (h *InviteHandler) AcceptInvite(c *fiber.Ctx) error {
	var req model.AcceptInviteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Call the service to accept the invite
	res, err := h.inviteService.AcceptInvite(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Invite accepted successfully", "data": res})
}
