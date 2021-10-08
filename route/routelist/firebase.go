package routelist

import (
	"code-echo/api/fcmfirebase"

	"github.com/labstack/echo/v4"
)

func GetNotif(echo *echo.Group) {

	//membuat method get untuk memanggil getNotifikasi
	echo.GET("/firebase", fcmfirebase.GetNotifikasiFCM)
	// echo.POST("/firebase", fcmfirebase.GetNotifikasiFCM)
}
