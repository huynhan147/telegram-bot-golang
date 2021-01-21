package main

import "telegram-bot/router"

func main() {
	r := router.New()
	r.Logger.Fatal(r.Start(":8888"))

}
