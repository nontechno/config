// Copyright 2020 NonTechno authors.

package config

import (
	log "github.com/sirupsen/logrus"
)

func GetConfigValue(key, fallback string) string {
	configInit()
	if len(defaultConfig) == 0 {
		log.Errorf("the default configuration is missing")
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

func GetConfigFlag(flag string, fallback bool) bool {
	configInit()
	if len(defaultConfig) == 0 {
		log.Errorf("the default configuration is missing")
		return fallback
	}
	key := flag + suffixConfigFlag
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