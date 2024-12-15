package feed

import (
	"api/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type req struct {
	Name     string `json:"name"`
	Contents string `json:"contents"`
}

func PostFeed(c echo.Context) error {
	b := new(req)
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}
	return nil
}

func GetFeed(c echo.Context) error {
	b := new(model.User)
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}
	return nil
}
