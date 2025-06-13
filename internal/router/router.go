package router

import (
	"github.com/example/internal/app"
)

func SetupRoutes(app *app.App) {
	router := app.Fiber.Group("/api/v1")
	SetupUserRoutes(router, app)
	SetupOrgRoutes(router, app)
	SetupInviteRoutes(router, app)
	SetupClockRoutes(router, app)
	SetupLeaveRoutes(router, app)
}
