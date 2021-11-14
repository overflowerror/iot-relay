package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Server struct {
		Bind    string `json:"bind"`
		Timeout int    `json:"timeout"` // in seconds
	} `json:"server"`
	Client struct {
		Address  string `json:"address"`
		Username string `json:"username"`
		Password string `json:"password"`
		DB       string `json:"db"`
		Host     string `json:"host"`
	} `json:"client"`
}

func ReadConfig(filename string) (Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	var config Config

	err = json.Unmarshal(content, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
