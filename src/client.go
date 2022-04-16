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

// Client wraps the type: ssh.Client for the use in this program.
type Client struct {
	// ssh client for sending in login requests.
	client *goph.Client

	// auth credentials.
	auth *goph.Client
}
