// Copyright 2020 The NonTechno Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func loadDefaultConfiguration() (map[string]string, error) {
	if len(os.Args) == 0 {
		return nil, errors.New("unknown executable")
	}
	base := os.Args[0]

	// implicit location of the config file: is where the binary is
	filename := strings.TrimSuffix(base, filepath.Ext(base)) + extConfig

	if len(defaultConfigExplicitName) == 0 {
		// let's try the working directory first (if there was no explicit setting)
		if config, err := LoadConfiguration(workingDirectoryConfigFileName); err == nil {
			defaultConfig = config
			return config, err
		}
	} else {
		filename = defaultConfigExplicitName
	}

	if config, err := LoadConfiguration(filename); err != nil {
		return nil, err
	} else {
		defaultConfig = config

		if value, exist := defaultConfig["ignore.absent.config.values"+suffixConfigFlag]; exist {
			reportMissingValues = !string2bool(value, false)
		}

		return config, err
	}
}
