package client

import (
	"fmt"
	"iot-relay/internal/config"
	"iot-relay/internal/types"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func lineProtocol(request types.Request, config config.Config) string {
	var builder strings.Builder
	builder.WriteString(config.Client.Measurement)
	builder.WriteString(",host=")
	builder.WriteString(config.Client.Host)
	builder.WriteString(",ip=")
	builder.WriteString(request.IP)
	builder.WriteString(",id=")
	builder.WriteString(request.ID)
	builder.WriteString(",loc=")
	builder.WriteString(request.Location)
	builder.WriteString(" ")

	for key, value := range request.Data {
		builder.WriteString(key)
		builder.WriteString("=")
		builder.WriteString(value)
		builder.WriteString(" ")
	}

	builder.WriteString(strconv.FormatInt(time.Now().UnixNano(), 10))

	return builder.String()
}

func GetHandler(config config.Config) types.Callback {
	return func(request types.Request) error {
		line := lineProtocol(request, config)

		url := config.Client.Address + "/write?db=" + config.Client.DB

		resp, err := http.Post(url, "application/octet-stream", strings.NewReader(line))
		if err != nil {
			log.Println(err)
			return err
		}
		if resp.StatusCode != 204 {
			err = fmt.Errorf("influxdb responded with %d", resp.StatusCode)
			log.Println(err)
			return err
		}

		return nil
	}
}
