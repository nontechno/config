// Copyright 2020 The NonTechno Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"sort"
	"strings"
)

func onKeyFound(key string) {
	if key == suffixConfigFlag {
		key = suffixConfigFlag
	}

	defaultConfigGuard.Lock()
	defaultConfigStats[key]++
	defaultConfigGuard.Unlock()
}

func onKeyMissing(key string) {
	onKeyFound(key)
}

// this func produces textual report of the usage pattern of the defalt confuguration
func UsageReport() string {
	defaultConfigGuard.Lock()
	defer defaultConfigGuard.Unlock()

	const (
		askedAndFound    = 1
		askedAndMissing  = 2
		notAskedButFound = 3
	)

	meaning := map[int]string{
		askedAndFound:    "asked + found",
		askedAndMissing:  "asked - missing",
		notAskedButFound: "unused",
	}
	stats := map[string]int{}

	for key, _ := range defaultConfigStats {
		if _, found := defaultConfig[key]; found {
			stats[key] = askedAndFound
		} else {
			stats[key] = askedAndMissing
		}
	}

	for key, _ := range defaultConfig {
		if _, found := defaultConfigStats[key]; !found {
			stats[key] = notAskedButFound
		}
	}

	sortedKeys := make([]string, 0, len(stats))
	for key, _ := range stats {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	report := "Configuration report:\n"
	for _, key := range sortedKeys {
		value := stats[key]
		if key != keyConfigurationFileName && !strings.HasPrefix(key, keyDebugPrefix) {
			report += fmt.Sprintf("    [%s] = %s\n", key, meaning[value])
		}
	}
	return report
}

var (
	ConfigUsageReport = UsageReport
)
