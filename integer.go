// Copyright 2020 The NonTechno Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"strconv"
	"strings"
)

func GetInt(key string, min, max, fallback int) int {
	value := strings.Trim(GetValue(key, ""), trimCharset)
	if len(value) == 0 {
		return fallback
	}

	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		// todo: report the error maybe?
		return fallback
	}

	if int64(int(parsed)) != parsed {
		return fallback
	}

	result := int(parsed)

	if result < min || result > max {
		return fallback
	}
	return result
}
