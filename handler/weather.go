package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"telegram-bot/config"
	"telegram-bot/models"
)

func getWeather() string {
	config, err := config.LoadConfig("./")

	if err != nil {
		log.Fatal("can not load config: ", err)
	}

	response, err := http.Get(config.APIUrlWeather + "?id=1581130" + "&appid=" + config.APIKeyWeather + "&&lang=vi")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var dataResponseWeaher models.ResponseWeather
	// text := string(body)

	err1 := json.Unmarshal([]byte(body), &dataResponseWeaher)
	// fmt.Println(body)
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Printf("value: %+v", dataResponseWeaher)
	temperature := strconv.FormatFloat((dataResponseWeaher.Main.Temp - 273.15), 'f', 2, 64)
	description := dataResponseWeaher.Weather[0].Description
	nameLocation := dataResponseWeaher.Name
	return "Địa điểm: " + nameLocation + ", Nhiệt độ: " + temperature + " độ C, Mô tả: " + description

}
