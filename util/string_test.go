package util

import "testing"

func TestNonEmptyOrDefult(t *testing.T) {
	def := "default"

	if def != NonEmptyOrDefult("", def) {
		t.Errorf("failed to return default")
	}

	if NonEmptyOrDefult("ok", def) != "ok" {
		t.Errorf("failed to return original")
	}
}
