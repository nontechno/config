// Copyright 2020 The NonTechno Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"strings"
	"time"
)

func GetDuration(key string, min, max, fallback time.Duration) time.Duration {
	value := strings.Trim(GetValue(key, ""), " \t\r\n")
	if len(value) == 0 {
		return fallback
	}
	result, err := time.ParseDuration(value)
	if err != nil {
		// todo: report the error maybe?
		return fallback
	}

	if result < min || result > max {
		return fallback
	}
	return result
}
