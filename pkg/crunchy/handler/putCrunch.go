package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/tmli3b3rm4n/crunchyRoll/pkg/crunchy/model"
	"net/http"
)

func PutCrunch(c echo.Context) error {
	crunch := model.Crunchy{}
	if err := c.Bind(&crunch).Error(); err != "" {
		return c.JSON(http.StatusBadRequest, err)
	}
	// We would update record here.

	return c.JSON(http.StatusOK, crunch)
}
