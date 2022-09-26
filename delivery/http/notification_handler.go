package http

import (
	"notif-engine/common"
	"notif-engine/model"
	notifSrv "notif-engine/service"

	"github.com/labstack/echo/v4"
)

func NewNotificationHandler(app *echo.Group, notifSrv notifSrv.PublishService) {
	app.POST("/", PublishNotif(notifSrv))
	app.POST("/artikel", PublishNotificationArtikel(notifSrv))
	app.POST("/kabar-donasi", PublishNotificationKabarDonasi(notifSrv))
}

func PublishNotif(notifSrv notifSrv.PublishService) echo.HandlerFunc {
	return func(c echo.Context) error {
		
		payload := &model.PayloadNotificationRequest{}
		c.Bind(payload)

		if err := common.ValidateStruct(payload); err != nil {
			return c.JSON(400, err)
		}

		// call service	
		result, err := notifSrv.PublishNotif(c.Request().Context(), payload)
		if err != nil {
			return c.JSON(400, err)
		}

		return common.SuccessResponse(c, 200, result)
	}
}

func PublishNotificationArtikel(notifSrv notifSrv.PublishService) echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := &model.PayloadNotificationArtikel{}
		c.Bind(payload)

		if err := common.ValidateStruct(payload); err != nil {
			return c.JSON(400, err)
		}

		// call service	
		result, err := notifSrv.PublishNotificationArtikel(c.Request().Context(), payload)
		if err != nil {
			return c.JSON(400, err)
		}

		return common.SuccessResponse(c, 200, result)
	}
}

func PublishNotificationKabarDonasi(notifSrv notifSrv.PublishService) echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := &model.PayloadNotificationKabarDonasi{}
		c.Bind(payload)

		if err := common.ValidateStruct(payload); err != nil {
			return c.JSON(400, err)
		}

		// call service	
		result, err := notifSrv.PublishNotificationKabarDonasi(c.Request().Context(), payload)
		if err != nil {
			return c.JSON(400, err)
		}

		return common.SuccessResponse(c, 200, result)
	}
}

