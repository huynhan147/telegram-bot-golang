package config

import "github.com/spf13/viper"

// Config : config app
type Config struct {
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`

	TokenAPIBotTelegram string `mapstructure:"TokenAPIBotTelegram"`
	URLAPIBotTelegram   string `mapstructure:"URLAPIBotTelegram"`
}

// LoadConfig : load config from env
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// const (
// 	// Config Mysql DB
// 	USER     string = ""
// 	PASSWORD string = ""
// 	DB       string = ""
// 	HOST     string = ""
// 	PORT     string = ""
// 	// Config telegram bot
// 	TokenAPIBotTelegram string = "1380062376:AAH8GQDHg9OJgJ3qux6FgfO1WWNrDQ60-Oo"
// 	URLAPIBotTelegram   string = "https://api.telegram.org/bot"
// )
