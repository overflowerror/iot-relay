package server

import (
	"bufio"
	"iot-relay/internal/types"
	"log"
	"net"
	"strings"
)

func connectionHandler(conn net.Conn) {
	reader := bufio.NewReader(conn)

	defer func() {
		log.Println("closing connection")
		_ = conn.Close()
	}()

	var request types.Request
	request.Data = make(map[string]string)
	request.IP = strings.Split(conn.RemoteAddr().String(), ":")[0]

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			log.Printf("connection error: %v", err)
			break
		}

		if len(line) == 1 {
			// end of request

			if !request.IsValid() {
				log.Println("protocol error")
				_, _ = conn.Write([]byte("bad\n"))
				break
			}

			err = callback(request)
			if err != nil {
				// should be logged by callback
				_, _ = conn.Write([]byte("fail\n"))
			} else {
				_, _ = conn.Write([]byte("ok\n"))
			}
			break
		} else {
			err = parseLine(line, &request)
			if err != nil {
				log.Printf("parsing error: %v", err)
				_, _ = conn.Write([]byte("bad\n"))
				break
			}
		}
	}
}
