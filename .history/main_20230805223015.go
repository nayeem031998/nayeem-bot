package main

import (
	"fmt"

	"nayeem/bot"
	"github.com/akhil/discord-ping/config"
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