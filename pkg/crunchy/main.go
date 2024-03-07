package main

import (
	"github.com/labstack/echo/v4"
	"github.com/tmli3b3rm4n/crunchyRoll/pkg/crunchy/routes"
)

func main() {
	e := new(echo.Echo)
	routes.RegisterRoutes(e)
	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
