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

func SetupOrgRoutes(router fiber.Router, app *app.App) {
	orgRoutes := router.Group("/orgs", middleware.AuthMiddleware())
	orgRepo := repository.NewRepository[entity.Organization](app.DB)
	inviteRepo := repository.NewRepository[entity.Invite](app.DB)
	service := service.NewOrgService(*orgRepo, *inviteRepo)
	handler := handler.NewOrgHandler(service)

	orgRoutes.Post("/createOrg", handler.CreateOrg)
	orgRoutes.Post("/:orgId/inviteEmployee", handler.InviteEmployee)
}
