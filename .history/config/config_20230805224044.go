package config

import (
	"encoding/json"
	"fmt"
	"io"
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

	file, err := io

	if err != nil {
		fmt.Println((err.Error()))
		return err
	}

	fmt.Println(file)

	err = json.NewDecoder(file).Decode(&config)


	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}