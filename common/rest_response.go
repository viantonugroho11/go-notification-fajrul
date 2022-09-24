package common

import "github.com/labstack/echo/v4"

func SuccessResponse(c echo.Context, status int, data interface{}) error {
	return c.JSON(status, data)
}