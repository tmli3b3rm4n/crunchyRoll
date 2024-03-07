package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/tmli3b3rm4n/crunchyRoll/pkg/crunchy/model"
	"net/http"
)

func GetCrunchy(c echo.Context) error {
	crunchy := model.Crunchy{Name: c.Param("name")}

	return c.JSON(http.StatusOK, crunchy)
}
