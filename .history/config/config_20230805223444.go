package config

import (
	"encoding/json"
	"fmt"
	"os"
	//"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := os.Open("./config.json")

	if err != nil {
		fmt.Println((err.Error()))
		return err
	}

	fmt.Println(file)

	err = 

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}