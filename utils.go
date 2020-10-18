// Copyright 2020 NonTechno authors.

package config

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func expand(what string, retainValue bool) string {
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
