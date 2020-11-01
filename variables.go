// Copyright 2020 The NonTechno Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import "sync"

var (
	defaultConfig             map[string]string
	defaultConfigGuard        sync.Mutex
	defaultConfigExplicitName string
	defaultConfigStats        map[string]int
	reportMissingValues       = true
)
