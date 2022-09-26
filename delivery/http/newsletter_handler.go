package http

import (
	"notif-engine/common"
	newsletterSrv "notif-engine/service"

	"github.com/labstack/echo/v4"
)


func NewNewsletterHandler(app *echo.Group, newsSrv newsletterSrv.NewsletterService) {
	app.GET("/", GetNewsAll(newsSrv))
}

func GetNewsAll(newsSrv newsletterSrv.NewsletterService) echo.HandlerFunc {
	return func(c echo.Context) error {
		// call service	
		result, err := newsSrv.GetAllNewsletter(c.Request().Context())
		if err != nil {
			return c.JSON(400, err)
		}

		return common.SuccessResponse(c, 200, result)
	}
}
