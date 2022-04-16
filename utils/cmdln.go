package utils

/*
// Author: Aliasgar Khimani (NovusEdge)
// Project: github.com/NovusEdge/buraq
//
// Copyright: GNU General Public License v3.0
// See the LICENSE file for more info.
//
// All Rights Reserved
*/

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// GetHomeDirectory returns the path to user's home directory ($HOME)
func GetHomeDirectory() string {
	return os.Getenv("HOME")
}

// GetEnv fetches the env for buraq
// The default location for env is ~/.buraq/env
func GetEnv() map[string]string {
	homeDir := GetHomeDirectory()
	content, err := ioutil.ReadFile(homeDir + "/.buraq/env")
	if err != nil {
		log.Fatal(err)
	}

	env := make(map[string]string)
	splitContent := strings.Split(string(content), "\n")
	splitContent = splitContent[:len(splitContent)-1]
	for _, s := range splitContent {
		tmp := strings.Split(s, "=")
		if len(tmp) == 1 {
			env[tmp[0]] = ""
		} else {
			env[tmp[0]] = tmp[1]
		}
	}

	return env
}
