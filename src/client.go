package gydeon

/*
// Author: Aliasgar Khimani (NovusEdge)
// Project: github.com/NovusEdge/gydeon
//
// Copyright: GNU General Public License v3.0
// See the LICENSE file for more info.
//
// All Rights Reserved
*/

import (
	"golang.org/x/crypto/ssh"
)

// Client wraps the type: ssh.Client for the use in this program.
type Client struct {
	// sshclient
	sshclient ssh.Client
}
