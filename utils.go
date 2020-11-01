// Copyright 2020 The NonTechno Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"os"
	"strings"

	str "github.com/nontechno/strings"
	log "github.com/sirupsen/logrus"
)

func expand(what string, retainValue bool) string {

	const prefix = "${"
	const postfix = "}"

	resolver := str.EnvironmentResolver
	if retainValue {
		resolver = str.EnvironmentResolverIntact
	}

	logger := log.WithField("domain", "config.expand")
	result, err := str.Expand(what, prefix, postfix, resolver, logger)
	if err != nil {
		logger.WithError(err).Errorf("unexpected failure during string expansion")
	}
	return result
}

func expandAlt(what string, retainValue bool) string {
	return os.Expand(what, func(what string) string {
		if expanded, found := os.LookupEnv(what); found {
			return expanded
		} else {
			// more "advanced" means of reporting may not be available at this point (in startup sequence)
			log.Warningf("Environment variable (%s) not found !!!", what)
			if retainValue {
				return fmt.Sprintf("${%s}", what)
			} else {
				return ""
			}
		}
	})
}

func string2bool(value string, fallback bool) bool {
	switch strings.ToLower(value) {
	case "yes", "affirmative", "true":
		return true
	case "no", "false", "nada":
		return false
	case "":
		return fallback
	default:
		return fallback
	}
}
