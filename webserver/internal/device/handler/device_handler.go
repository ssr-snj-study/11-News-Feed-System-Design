package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"webserver/internal/device/entity"
	"webserver/internal/device/service"
)

type DeviceHandler struct {
	DeviceService *service.DeviceService
}

func (h *DeviceHandler) Device(c echo.Context) error {
	req := new(entity.Req)
	if err := c.Bind(req); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	userId, err := h.DeviceService.UpsertDevice(req)
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"userId": userId,
	}

	return c.JSON(http.StatusOK, response)
}
