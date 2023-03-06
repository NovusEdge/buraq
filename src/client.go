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
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

// SSHConnect tries to connect to the ssh service with given credentials
func SSHConnect(username, host, password string, port uint16, timeout time.Duration) (bool, error) {
	sshConfig := &ssh.ClientConfig{
		User:    username,
		Auth:    []ssh.AuthMethod{ssh.Password(password)},
		Timeout: timeout,
	}

	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	service := fmt.Sprintf("%s:%d", host, port)

	client, clientErr := ssh.Dial("tcp", service, sshConfig)
	if clientErr != nil {
		return false, clientErr
	}

	// Close the client connection at the end of attempt...
	defer client.Close()

	session, sessionErr := client.NewSession()
	if sessionErr != nil {
		return false, sessionErr
	}
	defer session.Close()

	return true, nil
}
