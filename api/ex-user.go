package api

import (
	"code-echo/db"
	"code-echo/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	db := db.Manager()
	user := []model.User{}
	rows := db.Limit(3).Find(&user)

	return c.JSON(http.StatusOK, rows)
}
