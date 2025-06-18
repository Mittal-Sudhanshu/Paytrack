package router

import (
	"github.com/example/internal/app"
	"github.com/example/internal/entity"
	"github.com/example/internal/handler"
	"github.com/example/internal/middleware"
	"github.com/example/internal/repository"
	"github.com/example/internal/service"
	"github.com/gofiber/fiber/v2"
)

func SetupPayrollRoutes(router fiber.Router, app *app.App) {
	payrollRepo := repository.NewRepository[entity.Payroll](app.DB)
	clockRepo := repository.NewRepository[entity.ClockEntry](app.DB)
	leaveRequestRepo := repository.NewRepository[entity.LeaveRequest](app.DB)
	employeeRepo := repository.NewRepository[entity.Employee](app.DB)
	leaveBalaceRepo := repository.NewRepository[entity.LeaveBalance](app.DB)
	payrollService := service.NewPayrollService(*payrollRepo, *clockRepo, *leaveRequestRepo, *employeeRepo, *leaveBalaceRepo)
	payrollHandler := handler.NewPayrollHandler(payrollService)
	payrollRoutes := router.Group("/payroll", middleware.AuthMiddleware())
	payrollRoutes.Get("/:orgId", payrollHandler.GetPayrolls)
	payrollRoutes.Post("/:employeeId/generate", payrollHandler.GeneratePayroll)
}
