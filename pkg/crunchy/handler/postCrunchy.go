package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/tmli3b3rm4n/crunchyRoll/pkg/crunchy/model"
	"net/http"
)

func PostCrunch(c echo.Context) error {
	crunch := model.Crunchy{}
	if err := c.Bind(&crunch).Error(); err != "" {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, crunch)
}
