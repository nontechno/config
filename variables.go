// Copyright 2020 NonTechno authors.

package config

import "sync"

var (
	defaultConfig             map[string]string
	defaultConfigGuard        sync.Mutex
	defaultConfigExplicitName string
	defaultConfigStats        map[string]int
	reportMissingValues       = true
)
