package config

import (
	"fmt"
	"testing"
)

func Test_Report(t *testing.T) {
	_ = GetConfigValue("key", "fallback")
	fmt.Println("", UsageReport())
}
