package server

import (
	"iot-relay/internal/config"
	"iot-relay/internal/types"
	"log"
	"net"
	"time"
)

var callback types.Callback

func Listen(config config.Config, _callback types.Callback) error {
	callback = _callback

	listener, err := net.Listen("tcp", config.Server.Bind)
	if err != nil {
		return err
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("connection from %v", connection.RemoteAddr())

			_ = connection.SetReadDeadline(time.Now().Add(time.Second * time.Duration(config.Server.Timeout)))
			go connectionHandler(connection)
		}
	}
}
