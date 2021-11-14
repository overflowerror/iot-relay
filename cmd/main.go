package main

import (
	. "iot-relay/internal/client"
	. "iot-relay/internal/config"
	. "iot-relay/internal/server"
	"log"
)

const configFile = "config.json"

func main() {
	log.Println("reading config file")
	config, err := ReadConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("starting server")
	err = Listen(config, GetHandler(config))
	if err != nil {
		log.Fatalln(err)
	}
}
