package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"telegram-bot/config"

	"github.com/labstack/echo/v4"
)

type ResponseTeleGram struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
	ErrorCode   int    `json:"error_code"`
}

// Send : send message
func Send(text string, chatID int64) {
	config, err := config.LoadConfig("./")
	urlAPI := config.URLAPIBotTelegram + config.TokenAPIBotTelegram + "/sendMessage"
	if err != nil {
		log.Fatal("can not load config: ", err)
	}
	strChatID := strconv.FormatInt(chatID, 10)

	response, err := http.PostForm(urlAPI, url.Values{
		"method":     {"sendMessage"},
		"chat_id":    {strChatID},
		"text":       {text},
		"parse_mode": {"HTML"}})

	if err != nil {
		log.Fatal("Request Fail ", err)
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var dataResponse ResponseTeleGram
	// text := string(body)

	err1 := json.Unmarshal([]byte(body), &dataResponse)
	// fmt.Println(body)
	if err1 != nil {
		log.Fatal(err1)
	}

	if !dataResponse.Ok {
		fmt.Printf("Error : %+v\n", dataResponse.Description)
	}

	// var res map[string]interface{}

	// json.NewDecoder(resp.Body).Decode(&res)

	// // fmt.Println(res["json"])
	// body, _ := ioutil.ReadAll(req.Body)

	// // resp := ResponseTeleGram{}
	// // err = json.Unmarshal(body, &resp)
	log.Printf("value: %+v", dataResponse)
}

func TestSend(c echo.Context) error {
	Send("huytest", 848442722)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"statusCode": 200,
		"message":    "xong",
	})

}
