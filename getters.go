// Copyright 2020 The NonTechno Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	log "github.com/sirupsen/logrus"
)

func GetValue(key, fallback string) string {
	configInit()
	if len(defaultConfig) == 0 {
		log.Errorf("the default configuration is missing")
		onKeyMissing(key)
		return fallback
	}
	if value, present := defaultConfig[key]; present {
		onKeyFound(key)
		return value
	}
	onKeyMissing(key)

	if len(fallback) == 0 && reportMissingValues {
		// let's assume that if fallback value is missing - it is alarming that it is not present
		log.Errorf("the loaded configuration is missing value for key: (%s); config file (%s)", key, defaultConfig[keyConfigurationFileName])
	}
	return fallback
}

func GetFlag(flag string, fallback bool) bool {
	configInit()
	key := flag + suffixConfigFlag
	if len(defaultConfig) == 0 {
		log.Errorf("the default configuration is missing")
		onKeyMissing(key)
		return fallback
	}
	if value, present := defaultConfig[key]; present {
		onKeyFound(key)
		return string2bool(value, fallback)
	} else {
		onKeyMissing(key)
	}

	if reportMissingValues {
		log.Warningf("the loaded configuration is missing value for key: (%s); config file (%s)", key, defaultConfig[keyConfigurationFileName])
	}
	return fallback
}

var (
	GetConfigValue = GetValue
	GetConfigFlag  = GetFlag
)
