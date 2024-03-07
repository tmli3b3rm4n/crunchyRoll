package main

import (
	"github.com/labstack/echo/v4"
	"github.com/tmli3b3rm4n/crunchyRoll/pkg/crunchy/routes"
)

func main() {
	// Create an Echo instance
	e := echo.New()

	// Register routes
	routes.RegisterRoutes(e)

	// Start the server
	e.Start(":8080")
}
