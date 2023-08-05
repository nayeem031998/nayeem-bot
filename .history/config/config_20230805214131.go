package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token string
	BotPrefix string
	config *ConfigStruct
)

type ConfigStruct struct {

	Token string `json:"Token"`

	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {

	fmt.Println("Reading from config.json")

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err)
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
