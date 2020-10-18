package config

import (
	"testing"
)

func Test_GetConfigValue(t *testing.T) {
	value := GetConfigValue("key", "fallback")
	if value != "bingo" {
		t.Fatalf("failed config getter test")
	}
}
