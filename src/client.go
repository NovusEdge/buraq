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
	"log"
	"time"

	utils "github.com/NovusEdge/buraq/utils"
	"golang.org/x/crypto/ssh"
)

// AttemptConnection tries to connect to the ssh server using credentials
// provided.
func AttemptConnection(proto, username, host, password string, port uint, timeout time.Duration) (bool, error) {
	if proto != "tcp" || proto != "udp" {
		log.Println(utils.ColorIt(utils.ColorYellow, fmt.Sprintf("[W]: Did not recognize protocol: %s for ssh dialup.\nUsing tcp as default", proto)))
		proto = "tcp"
	}

	sshConfig := &ssh.ClientConfig{
		User:    username,
		Auth:    []ssh.AuthMethod{ssh.Password(password)},
		Timeout: timeout,
	}

	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, clientErr := ssh.Dial(proto, fmt.Sprintf("%s:%d", host, port), sshConfig)
	if clientErr != nil {
		return false, clientErr
	}

	defer client.Close()

	session, sessionErr := client.NewSession()
	if sessionErr != nil {
		return false, sessionErr
	}
	defer session.Close()

	_, err := session.Output("whoami")
	if err != nil {
		return false, err
	}

	return true, nil
}
