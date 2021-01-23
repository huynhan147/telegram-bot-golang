package models

type ResponseWeather struct {
	Main    Main      `json:"main"`
	Name    string    `json:"name"`
	Weather []Weather `json:"weather"`
}

type Main struct {
	Temp float64 `json:"temp"`
}

type Weather struct {
	Description string `json:"description"`
}
