package device

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	isError    = 0
	toBeCreate = 1
	toBeUpdate = 2
)

type params struct {
	Name   string `json:"name"`
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}

func UpsertDevice(c echo.Context) error {
	b := new(params)
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	flag, err := CheckDevice(b)
	if err != nil && b.UserId == 0 {
		return err
	}
	switch flag {
	case toBeCreate:
		err = createDevice(b)
		if err != nil {
			return err
		}
	case toBeUpdate:
		err = updateDevice(b)
		if err != nil {
			return err
		}
	}

	response := map[string]interface{}{
		"userId": b.UserId,
	}
	return c.JSON(http.StatusOK, response)
}
