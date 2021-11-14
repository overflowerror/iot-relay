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

func lineProtocolPrefix(request types.Request, config config.Config) string {
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

	return builder.String()
}

func GetHandler(config config.Config) types.Callback {
	return func(request types.Request) error {
		prefix := lineProtocolPrefix(request, config)
		timeString := strconv.FormatInt(time.Now().UnixNano(), 10)

		url := config.Client.Address + "/write?db=" + config.Client.DB

		client := &http.Client{}

		var builder strings.Builder

		for key, value := range request.Data {
			builder.Reset()
			builder.WriteString(prefix)

			builder.WriteString(key)
			builder.WriteString("=")
			builder.WriteString(value)
			builder.WriteString(" ")

			builder.WriteString(timeString)

			req, err := http.NewRequest("POST", url, strings.NewReader(builder.String()))

			if len(config.Client.Username) != 0 && len(config.Client.Password) != 0 {
				req.SetBasicAuth(config.Client.Username, config.Client.Password)
			}

			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
				return err
			}
			if resp.StatusCode != 204 {
				err = fmt.Errorf("influxdb responded with %d", resp.StatusCode)
				log.Println(err)
				return err
			}
		}

		return nil
	}
}
