package exceptions

import (
	"echo-boilerplate/utils/helpers"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	InternalServerErrorMsg = "internal server error"
	ForbiddenErrorMsg      = "forbidden error"
	NotFoundErrorMsg       = "not found error"
	BadRequestErrorMsg     = "bad request error"
	ValidationErrorMsg     = "validation error"
)

func ExceptionHandler(err error, c echo.Context) {
	if NotFoundError(err, c) {
		return
	} else if ForbiddenError(err, c) {
		return
	} else if BadRequestError(err, c) {
		return
	} else if ValidationError(err, c) {
		return
	} else {
		InternalServerError(err, c)
		return
	}
}

func InternalServerError(err error, c echo.Context) bool {

	if err.Error() == InternalServerErrorMsg {
		c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
		return true
	}

	return false
}

func NotFoundError(err error, c echo.Context) bool {

	if err.Error() == NotFoundErrorMsg {
		c.JSON(http.StatusNotFound, helpers.FailedResponse(err.Error()))
		return true
	}

	return false
}

func ForbiddenError(err error, c echo.Context) bool {

	if err.Error() == ForbiddenErrorMsg {
		c.JSON(http.StatusForbidden, helpers.FailedResponse(err.Error()))
		return true
	}

	return false
}

func BadRequestError(err error, c echo.Context) bool {

	if err.Error() == BadRequestErrorMsg {
		c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		return true
	}

	return false
}

func ValidationError(err error, c echo.Context) bool {
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		var report string
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report = fmt.Sprintf("%s is required",
					err.Field())
			case "email":
				report = fmt.Sprintf("%s is not valid email",
					err.Field())
			case "gte":
				report = fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param())
			case "lte":
				report = fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param())
			}

			break
		}
		c.JSON(http.StatusBadRequest, helpers.FailedResponse(report))
		return true
	}
	return false
}
