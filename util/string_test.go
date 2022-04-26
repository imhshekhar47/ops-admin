package util

import "testing"

func TestEncode(t *testing.T) {
	key := "abcd.123@"
	coded := Encode(key)
	if coded != Encode(key) {
		t.Errorf("encoding is not consistent")
		t.Fail()
	}
}

func TestNonEmptyOrDefult(t *testing.T) {
	def := "default"

	if def != NonEmptyOrDefult("", def) {
		t.Errorf("failed to return default")
	}

	if NonEmptyOrDefult("ok", def) != "ok" {
		t.Errorf("failed to return original")
	}
}
