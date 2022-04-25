package util

import "strings"

func NonEmptyOrDefult(value string, def string) string {
	if len(strings.Trim(value, " ")) == 0 {
		return def
	} else {
		return value
	}
}
