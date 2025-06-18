package handler

import (
	"time"

	"github.com/example/internal/service"
	"github.com/gofiber/fiber/v2"
)

type PayrollHandler struct {
	payrollService service.PayrollService
}

func NewPayrollHandler(payrollService service.PayrollService) *PayrollHandler {
	return &PayrollHandler{
		payrollService: payrollService,
	}
}

func (h *PayrollHandler) GetPayrolls(c *fiber.Ctx) error {
	// orgId := c.Params("orgId")
	// // payroll, error := h.payrollService.GeneratePayroll(c.Context(), orgId)
	// if error != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, error.Error())
	// }
	// return c.JSON(payroll)
	return nil
}

func (h *PayrollHandler) GeneratePayroll(c *fiber.Ctx) error {
	employeeId := c.Params("employeeId") // or get from token
	type Request struct {
		Month int `json:"month"`
		Year  int `json:"year"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid input")
	}

	payroll, err := h.payrollService.GeneratePayroll(c.Context(), employeeId, time.Month(req.Month), req.Year)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"data": payroll,
	})
}
