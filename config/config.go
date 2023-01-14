// Copyright (C) Oleg Lysiak - All Rights Reserved
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	// json token and bot prefix

	Token     string `json : "MTAzOTg5NjMxNDk2MTQ3NzY3Mg.Ge1Zpi.k3lsJUBf3UtyoZr8ey2lG8n3w2mW_S3bJmfHQo"`
	BotPrefix string `json : "$$"`
}

// reading of .json file

func ReadConfig() error {
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
