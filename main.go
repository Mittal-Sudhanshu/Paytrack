// main.go
package main

import (
	"github.com/example/internal/app"
	"github.com/example/internal/router"
)

func main() {
	app, _ := app.InitializeApp()
	router.SetupRoutes(app)
	app.Start()
}
