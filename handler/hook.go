package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"telegram-bot/models"

	"github.com/labstack/echo/v4"
)

// WebHook : webhook telegram
func WebHook(c echo.Context) error {
	message := models.Response{}
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("fail: %s", err)
		return c.String(http.StatusBadRequest, "errr")
	}

	err = json.Unmarshal(b, &message)
	if err != nil {
		log.Printf("fail: %s", err)
		return c.String(http.StatusBadRequest, "errr")
	}

	log.Printf("value: %+v", message.Message)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"statusCode": 200,
		"message":    message.Message,
	})
}
