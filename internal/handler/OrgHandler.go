package handler

import (
	"fmt"

	"github.com/example/internal/model"
	"github.com/example/internal/service"
	"github.com/gofiber/fiber/v2"
)

type OrgHandler struct {
	orgService service.OrgService
}

func NewOrgHandler(orgService service.OrgService) *OrgHandler {
	return &OrgHandler{orgService: orgService}
}

func (h *OrgHandler) CreateOrg(c *fiber.Ctx) error {
	var req model.CreateOrgRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	data, err := h.orgService.CreateOrg(c.Context(), req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": data})
}

func (h *OrgHandler) InviteEmployee(c *fiber.Ctx) error {
	var req model.InviteEmployeeRequest

	if err := c.BodyParser(&req); err != nil {
		fmt.Println("Error parsing request body:", err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	orgId := c.Params("orgId")
	userId := c.Locals("user_id").(string)
	fmt.Println("User ID from context:", userId)
	if orgId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "orgId is required")
	}
	data, err := h.orgService.InviteEmployee(c.Context(), req, orgId, userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": data})
}
func (h *OrgHandler) GetMyOrgs(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)
	if userId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "user_id is required")
	}
	data, err := h.orgService.GetMyOrgs(c.Context(), userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": data})
}
