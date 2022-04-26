package util

import (
	"encoding/base64"
	"strings"
)

func Encode(key string) string {
	return base64.StdEncoding.EncodeToString([]byte(key))
}

func NonEmptyOrDefult(value string, def string) string {
	if len(strings.Trim(value, " ")) == 0 {
		return def
	} else {
		return value
	}
}
