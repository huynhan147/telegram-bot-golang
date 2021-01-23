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

	log.Printf("value: %+v", message)
	userNameSender := message.Message.Sender.FirstName + " " + message.Message.Sender.LastName
	textResponse := "Bot này không có thời gian tán gẫu với bạn đâu nhé. \n Hãy dùng command để nói chuyện /help"
	switch message.Message.Text {
	case "/start":
		textResponse = "Chào mừng bạn " + userNameSender + " đã đến với Super Bot. Hãy sử dụng lệnh /help để tìm hiểu về tính năng của Bot"
	case "/help":
		textResponse = "Sử dụng: \n /weather - Lấy thông tin về thời tiết. \n /corona - Để lấy thông tin về Corona \n /subscribe - Để nhận tin nhắn Bot mỗi ngày"
	case "/weather":
		textResponse = getWeather()
		// textResponse = "Tính năng thời tiết đang trong quá trinh hoàn thiện"
	case "/donate":
		textResponse = "Ủng hộ tác giả tại VP Bank :193528139 \n Chân thành cảm ơn !!!"
	case "/corona":
		textResponse = "Tính năng cập nhật thông tin về corona đang trong quá trình hoàn thiện"

	}

	Send(textResponse, message.Message.Sender.ID)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"statusCode": 200,
		"message":    message.Message,
	})
}
