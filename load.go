// Copyright 2020 The NonTechno Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

func LoadConfiguration(filename string) (map[string]string, error) {
	filename = expand(filename, false)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		// log.Errorf("failed to open file (%s), due to (%v)", filename, err)
		return nil, err
	}

	// unify "new lines"
	txt := strings.ReplaceAll(string(bytes), "\r\n", "\n")
	txt = strings.ReplaceAll(txt, "\n\r", "\n")

	// cut off the tail
	if pos := strings.Index(txt, "\n..."); pos > 0 {
		txt = txt[:pos]
	}

	txt = expand(txt, false)

	// warning:
	//if there are multiple "segments" in this (yaml) file (separated by "\n---") -
	//we are going to read only the first one and discard the rest.

	raw := make(map[string]interface{})
	if err = yaml.Unmarshal([]byte(txt), &raw); err != nil {
		log.Errorf("failed to unmarshal content of file (%s), due to (%v)", filename, err)
		return nil, err
	}

	results := make(map[string]string)
	for k, v := range raw {
		if v == nil {
			// results[k] = ""
		} else if str, converts := v.(string); converts {
			results[k] = str
		} else if i, converts := v.(int); converts {
			results[k] = strconv.Itoa(i)
		} else if b, converts := v.(bool); converts {
			results[k] = strconv.FormatBool(b)
		} else if f, converts := v.(float64); converts {
			results[k] = strconv.FormatFloat(f, 'g', -1, 32)
		} else {
			log.Errorf("failed to ingest value of the key (%s) from file (%s)", k, filename)
		}
	}

	// let's remember where we loaded the config values from ...
	if abs, err := filepath.Abs(filename); err == nil {
		results[keyConfigurationFileName] = abs
	} else {
		results[keyConfigurationFileName] = filename
	}

	return results, nil
}
