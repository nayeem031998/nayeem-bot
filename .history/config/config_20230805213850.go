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
