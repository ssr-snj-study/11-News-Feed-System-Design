package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"webserver/internal/auth/entity"
	"webserver/internal/auth/service"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func (h *AuthHandler) Auth(c echo.Context) error {
	req := new(entity.Req)
	if err := c.Bind(req); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}
	cookie, err := h.AuthService.Authenticate(req)
	if err != nil {
		c.Logger().Warn(err)
	}
	c.SetCookie(cookie)

	return nil
}
