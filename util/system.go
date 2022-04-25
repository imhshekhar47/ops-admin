package util

import (
	"encoding/base64"
	"os"
)

func GetHostname() string {
	hostname, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	return hostname
}

func Encode(key string) string {
	return base64.StdEncoding.EncodeToString([]byte(key))
}
