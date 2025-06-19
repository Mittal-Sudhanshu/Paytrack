package router

import (
	"github.com/example/internal/app"
	"github.com/example/internal/entity"
	"github.com/example/internal/handler"
	"github.com/example/internal/repository"
	"github.com/example/internal/service"
	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes configures user-related routes
func SetupUserRoutes(router fiber.Router, app *app.App) {
	userRepo := repository.NewRepository[entity.User](app.DB)
	userService := service.NewUserService(*userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Public routes
	// fmt.Print(userRepo)
	users := router.Group("/users")
	users.Post("/register", userHandler.SignUp)
	users.Get("", userHandler.GetAllUsers)
	users.Put("/login", userHandler.Login)

	// Protected routes (require authentication)
	// protected := users.Group("/", middleware.AuthRequired())
	// protected.Get("/profile", userHandler.GetProfile)
	// protected.Put("/profile", userHandler.UpdateProfile)
	// protected.Delete("/profile", userHandler.DeleteProfile)
	// protected.Get("/", middleware.AdminRequired(), userHandler.GetAllUsers)
}
