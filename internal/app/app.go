// internal/app/app.go
package app

import (
	"log"
	"os"

	"github.com/example/internal/db"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type App struct {
	Fiber *fiber.App
	DB    *gorm.DB
}

func InitializeApp() (*App, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Create Fiber app with custom configuration
	fiberApp := fiber.New(fiber.Config{
		// Custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Default 500 status code
			code := fiber.StatusInternalServerError

			// Check if it's a fiber error and get the status code
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Log the error
			log.Printf("Error: %v", err)

			// Return JSON error response
			return c.Status(code).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		},
		// Enable case sensitive routing
		CaseSensitive: true,
		// Enable strict routing
		StrictRouting: true,
		// Server header
		ServerHeader: "Fiber",
		// App name
		AppName: "Example App v1.0.0",
	})

	// Add middleware
	fiberApp.Use(recover.New()) // Recover from panics
	fiberApp.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./internal/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}

	fiberApp.Use(swagger.New(cfg))
	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Register routers with dependencies
	// Example:
	// router.SetupOrgRoutes(fiberApp, app)
	// router.SetupUserRoutes(fiberApp, app)

	// Connect to DB
	database, err := db.ConnectDB()
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return nil, err
	}

	return &App{
		Fiber: fiberApp,
		DB:    database,
	}, nil
}

// Start starts the Fiber server
func (a *App) Start() error {
	// Get port from environment or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server starting on port %s", port)
	return a.Fiber.Listen(":" + port)
}

// Shutdown gracefully shuts down the server
func (a *App) Shutdown() error {
	log.Println("🔥 Shutting down server...")
	return a.Fiber.Shutdown()
}
