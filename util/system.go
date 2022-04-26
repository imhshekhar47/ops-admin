package util

import (
	"os"
)

func GetHostname() string {
	hostname, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	return hostname
}
