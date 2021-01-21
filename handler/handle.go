package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"telegram-bot/config"
	"telegram-bot/models"

	"github.com/labstack/echo/v4"
)

type dataGetMe struct {
	Ok     bool `json:"ok"`
	Result models.BotInfo
}

// Hello :  function abc
func Hello(e echo.Context) error {
	return e.JSON(200, map[string]interface{}{
		"statusCode": 200,
		"message":    "success ",
	})

}

// GetMe : get info bot
func GetMe(e echo.Context) error {

	config, err := config.LoadConfig("./")

	if err != nil {
		log.Fatal("can not load config: ", err)
	}

	response, err := http.Get(config.URLAPIBotTelegram + config.TokenAPIBotTelegram + "/getMe")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var dataResponse dataGetMe
	// text := string(body)

	err1 := json.Unmarshal([]byte(body), &dataResponse)
	// fmt.Println(body)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("%+v\n", dataResponse.Result)

	// Copy data from the response to standard output
	return e.JSON(200, map[string]interface{}{
		"statusCode": 200,
		"message":    dataResponse.Result,
	})
}
