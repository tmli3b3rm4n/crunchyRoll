package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func DeleteCrunch(c echo.Context) error {
	// We would update recode with delete timestamp.
	return c.JSON(http.StatusNoContent, "")
}
