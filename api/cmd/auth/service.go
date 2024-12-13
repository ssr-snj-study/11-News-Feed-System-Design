package auth

import (
	"api/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Auth(c echo.Context) error {
	b := new(model.User)
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}
	return nil
}
