package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/tmli3b3rm4n/crunchyRoll/pkg/crunchy/handler"
)

func RegisterRoutes(e *echo.Echo) {
	crunchy := e.Group("crunchy")
	crunchy.GET("/:name", handler.GetCrunchy)
	crunchy.DELETE("/:name", handler.DeleteCrunch)
	crunchy.PUT("/:name", handler.PutCrunch)
	crunchy.POST("/", handler.PostCrunch)
}
