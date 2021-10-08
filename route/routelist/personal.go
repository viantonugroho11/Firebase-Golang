package routelist

import (
	// "code-echo/api/firebase"
	"code-echo/api/peremajaan"

	"github.com/labstack/echo/v4"
)

func EventInstansi(echo *echo.Group) {

	echo.POST("/add", peremajaan.SimpanEvent)
	echo.POST("/update", peremajaan.UpdateEvent)
	echo.GET("/get", peremajaan.GetEvent)
}
