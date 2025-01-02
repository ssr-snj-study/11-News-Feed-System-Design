package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"webserver/internal/feed/entity"
	"webserver/internal/feed/service"
)

type FeedHandler struct {
	FeedService *service.FeedService
}

func (h *FeedHandler) PostFeed(c echo.Context) error {
	req := new(entity.Req)
	if err := c.Bind(req); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	h.FeedService.PostingFeed(req)

	return nil
}

func (h *FeedHandler) GetFeed(c echo.Context) error {
	req := new(entity.Req)
	if err := c.Bind(req); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	h.FeedService.Repo.GetFeed(req)

	return nil
}
