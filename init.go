// Copyright 2020 The NonTechno Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"os"
)

func configInit() {
	if defaultConfig != nil {
		return
	}
	defaultConfigGuard.Lock()
	defer defaultConfigGuard.Unlock()

	if defaultConfig != nil {
		// oh heey, looks like the config was loaded in a parallel thread
		return
	}
	loadDefaultConfiguration()
}

func init() {
	// use "on demand" config loading through configInit() call instead of using "init" func

	defaultConfigGuard.Lock()
	defer defaultConfigGuard.Unlock()

	if envFilename, found := os.LookupEnv(envConfigurationFileName); found && len(envFilename) > 0 {
		defaultConfigExplicitName = envFilename
	}

	// let's look at cli arguments and see of the config name was specified there
	for i, arg := range os.Args {
		if arg == argConfig && (i+1) < len(os.Args) {
			// found explicit location of the config file, specified in the command line
			defaultConfigExplicitName = os.Args[i+1]

			// let's remove these two args from the list of args
			os.Args = append(os.Args[:i], os.Args[i+2:]...)
			break
		}
	}

	defaultConfigStats = make(map[string]int)
}
