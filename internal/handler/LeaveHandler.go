package handler

import (
	"github.com/example/internal/model"
	"github.com/example/internal/service"
	"github.com/gofiber/fiber/v2"
)

type LeaveHandler struct {
	leaveService service.LeaveService // Assuming you have a LeaveService interface defined
}

func NewLeaveHandler(leaveService service.LeaveService) *LeaveHandler {
	return &LeaveHandler{
		leaveService: leaveService,
	}
}

func (h *LeaveHandler) ApplyLeave(c *fiber.Ctx) error {
	var leaveRequest model.LeaveRequest
	if err := c.BodyParser(&leaveRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	userId := c.Locals("user_id").(string) // Assuming user_id is set in middleware
	res, err := h.leaveService.ApplyLeave(c.Context(), userId, leaveRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to apply leave",
		})
	}

	return c.JSON(res)
}

func (h *LeaveHandler) GetLeaveBalance(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string) // Assuming user_id is set in middleware
	balance, err := h.leaveService.GetLeaveBalance(c.Context(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve leave balance",
		})
	}

	return c.JSON(balance)
}

func (h *LeaveHandler) GetLeaveRequests(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)
	requests, err := h.leaveService.GetLeaveRequests(c.Context(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve leave requests",
		})
	}
	return c.JSON(requests)
}

func (h *LeaveHandler) UpdateLeaveRequest(c *fiber.Ctx) error {
	requestId := c.Params("requestId")
	var status string
	if err := c.BodyParser(&status); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	if requestId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Request ID is required",
		})
	}

	userId := c.Locals("user_id").(string) // Assuming user_id is set in middleware
	_, err := h.leaveService.UpdateLeaveRequest(c.Context(), requestId, userId, status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to approve leave request",
		})
	}

	return c.JSON(fiber.Map{"message": "Leave request approved successfully"})
}
