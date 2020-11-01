// Copyright 2020 The NonTechno Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

const (
	extConfig                      = ".config"               // default extention of config file
	argConfig                      = "-config"               // cli argument to indicate the presence of config file name
	workingDirectoryConfigFileName = "./service" + extConfig // name of the config file to try to load from
	keyConfigurationFileName       = "config.filename"       // key fo the actual path/name of the loaded config file
	keyDebugPrefix                 = "debug."                // prefix of debug-related keys (not included in "report")
	envConfigurationFileName       = "CONFIG_FILENAME"       // env.var to specify the path/name of the config file to use
	suffixConfigFlag               = ".flag"                 // indicator of a boolean value
)
