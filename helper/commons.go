package helper

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)


func FailOnError(err error, msgerr string, msgsuc string) {
	if err != nil {
		log.Fatalf("%s: %s", msgerr, err)

	} else {
		fmt.Printf("%s\n", msgsuc)
	}

}

func CustomError(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required",
						err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email",
						err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s",
						err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s",
						err.Field(), err.Param())
				}

				break
			}
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}
}
func FindIn(what interface{}, where []interface{}) (idx int) {
	for i, v := range where {
		if v == what {
			return i
		}
	}
	return -1
}




