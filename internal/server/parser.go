package server

import (
	"errors"
	"iot-relay/internal/types"
	"strings"
)

var forbiddenKeys = []string{
	"ip",
}

func parseLine(line string, request *types.Request) error {
	token := strings.Split(line, "=")
	if len(token) != 2 {
		return errors.New("malformed; missing =")
	}

	key := strings.ToLower(token[0])
	value := strings.TrimRight(token[1], "\n")

	for _, forbidden := range forbiddenKeys {
		if forbidden == key {
			return errors.New("forbidden key")
		}
	}

	switch key {
	case "id":
		request.ID = value
	case "loc":
		request.Location = value
	default:
		request.Data[key] = value
	}

	return nil
}
