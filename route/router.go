package route

import (
	"code-echo/helper"
	"code-echo/route/routelist"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		// AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	e.Validator = &CustomValidator{validator: validator.New()}

	helper.CustomError(e)

	/*-------------------------------------------------
		LIST OF UNRESTRICTED ROUTE
		Put your unrestricted route function here with
		the imported package
	--------------------------------------------------*/

	/*-------------------------------------------------
		TOKEN VALIDATION WITH RSA256 ALGORITHM
		Validate JWT with public key from keycloak
		server. Put the public key path in config.json
		Do not change this code!!
	--------------------------------------------------*/
	g := e.Group("/event")
	f	:= e.Group("/firebase")
	/*-------------------------------------------------
		DISABLING JWT CUZ API GW ALREADY DID THAT
	key, _ := helper.GetRSAPublicKey()

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    key,
		SigningMethod: "RS256",
	}))
	---------------------------------------------------*/

	/*-------------------------------------------------
		LIST OF RESTRICTED ROUTE
		Put your restricted route function here with
		the imported package
	--------------------------------------------------*/

	routelist.EventInstansi(g)
	routelist.GetNotif(f)
	return e
}
