// Copyright 2020 NonTechno authors.

package config

import (
	"encoding/json"
	"errors"
	"strings"
)

func ExportConfiguration() ([]byte, error) {
	what := make(map[string]string)
	for key, value := range defaultConfig {
		lowercased := strings.ToLower(key)
		if strings.Contains(lowercased, "secret") || strings.Contains(lowercased, "password") {
			what[key] = strings.Repeat("*", len(value)) // remove confidential info
		} else {
			what[key] = value
		}
	}

	return json.Marshal(what)
}

func ImportConfiguration([]byte) error {
	return errors.New("ImportConfiguration: niy")
}
