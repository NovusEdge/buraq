package buraq

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
	"github.com/melbahja/goph"
)

// CheckCreds returns (true, nil) if connected to ssh service successfully.
func CheckCreds(username string, password string, target string, customConfig *goph.Config) (bool, error) {
	var err error
	var output []byte
	var c *goph.Client

	defer c.Close()

	if customConfig == nil {
		c, err = simpleConn(username, password, target)
	} else {
		c, err = customConn(customConfig)
	}

	if output, err = c.Run("whoami"); string(output) == c.Config.User {
		return true, err
	}

	return false, err
}

// simpleConn initiates a new ssh connection with standard config.
func simpleConn(username, password, target string) (c *goph.Client, err error) {
	auth := goph.Password(password)
	c, err = goph.New(username, target, auth)
	return
}


// customConn initiates a new ssh connection with custom config.
func customConn(conf *goph.Config) (*goph.Client, error) {
	return goph.NewConn(conf)
}
