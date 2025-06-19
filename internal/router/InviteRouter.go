package router

import (
	"github.com/example/internal/app"
	"github.com/example/internal/entity"
	"github.com/example/internal/handler"
	"github.com/example/internal/repository"
	"github.com/example/internal/service"
	"github.com/gofiber/fiber/v2"
)

func SetupInviteRoutes(router fiber.Router, app *app.App) {
	inviteRoutes := router.Group("/invites")

	// Here you would typically set up your invite-related handlers
	// For example:
	inviteRepo := repository.NewRepository[entity.Invite](app.DB)
	userRepo := repository.NewRepository[entity.User](app.DB)
	employeeRepo := repository.NewRepository[entity.Employee](app.DB)
	userOrgRepo := repository.NewRepository[entity.UserOrg](app.DB)
	userServiceImpl := service.NewUserService(*userRepo)
	inviteService := service.NewInviteService(*inviteRepo, *userRepo, *employeeRepo, *userOrgRepo, userServiceImpl)
	inviteHandler := handler.NewInviteHandler(inviteService)

	inviteRoutes.Post("/accept", inviteHandler.AcceptInvite)
	// inviteRoutes.Get("/:inviteId", inviteHandler.GetInvite)
	// inviteRoutes.Delete("/:inviteId", inviteHandler.DeleteInvite)

	// Note: The actual implementation of handlers and services is not shown here.
}
