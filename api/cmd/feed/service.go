package feed

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type req struct {
	Name     string `json:"name"`
	Contents string `json:"contents"`
	UserId   int    `json:"userId"`
	FeedId   int    `json:"feedId"`
}

func PostFeed(c echo.Context) error {
	fmt.Println("Headers:", c.Request().Header)
	b := new(req)
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	err := createFeed(b)
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	//getFollowerToken(b.UserId)
	//internal.SendMessageToFCM()

	return nil
}

func GetFeed(c echo.Context) error {
	b := new(req)
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	posting, err := getFeed(b)
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	response := map[string]interface{}{
		"contents": posting.Contents,
		"likes":    posting.Likes,
		"postDate": posting.CreatedTime,
	}
	return c.JSON(http.StatusOK, response)
}
