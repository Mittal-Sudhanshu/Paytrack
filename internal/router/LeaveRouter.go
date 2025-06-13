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

func SetupLeaveRoutes(router fiber.Router, app *app.App) {
	leaveRequestRepo := repository.NewRepository[entity.LeaveRequest](app.DB)
	leaveBalanceRepo := repository.NewRepository[entity.LeaveBalance](app.DB)
	leaveService := service.NewLeaveService(*leaveRequestRepo, *leaveBalanceRepo)
	leaveHandler := handler.NewLeaveHandler(leaveService)
	leaveRoutes := router.Group("/leaves", middleware.AuthMiddleware())
	leaveRoutes.Put("/apply", leaveHandler.ApplyLeave)
	leaveRoutes.Get("/balance", leaveHandler.GetLeaveBalance)
	leaveRoutes.Get("/requests", leaveHandler.GetLeaveRequests)
	leaveRoutes.Put("/:requestId", leaveHandler.UpdateLeaveRequest)
}
