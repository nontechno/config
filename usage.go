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
func ConfigUsageReport() string {
	defaultConfigGuard.Lock()
	defer defaultConfigGuard.Unlock()

	report := "Configuration report:\n"
	for key, count := range defaultConfigStats {
		report += fmt.Sprintf("    [%s] = %d\n", key, count)
	}

	// report present but unused config values
	unused := make([]string, 0, len(defaultConfig))
	for key, _ := range defaultConfig {
		if _, found := defaultConfigStats[key]; !found {
			if key != keyConfigurationFileName && !strings.HasPrefix(key, keyDebugPrefix) {
				unused = append(unused, key)
			}
		}
	}
	if len(unused) > 0 {
		sort.Strings(unused)
		report += "\nUnused configuration key(s):\n"
		for _, key := range unused {
			report += fmt.Sprintf("    [%s]\n", key)
		}
	}

	return report
}
