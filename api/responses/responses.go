package responses

import (
	"github.com/labstack/echo/v4"
	"key-value-store/api/models"
	"net/http"
)

func Err(c echo.Context, message string, code int) error {
	errors := map[string]interface{}{
		"error": message,
	}

	result := models.Response{
		Code:   code,
		Error:  true,
		Errors: errors,
		Result: nil,
	}

	return c.JSON(code, result)
}

func Err401(c echo.Context) error {
	errors := map[string]interface{}{
		"error": "You are not authorized for this operation",
	}

	result := models.Response{
		Code:   http.StatusUnauthorized,
		Error:  true,
		Errors: errors,
		Result: nil,
	}

	return c.JSON(http.StatusUnauthorized, result)
}

func Err500(c echo.Context, message string) error {
	errors := map[string]interface{}{
		"error": message,
	}

	result := models.Response{
		Code:   http.StatusInternalServerError,
		Error:  true,
		Errors: errors,
		Result: nil,
	}

	return c.JSON(http.StatusInternalServerError, result)
}

func Ok(c echo.Context, result interface{}) error {
	response := models.Response{
		Code:   http.StatusOK,
		Error:  false,
		Errors: nil,
		Result: result,
	}

	return c.JSON(http.StatusOK, response)
}
