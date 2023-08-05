package main

import (
	"fmt"

	"nayeem_bot/bot"
	"nayeem_bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}