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
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	utils "github.com/NovusEdge/buraq/utils"
)

// ENV is the environment variables for the program
var ENV = utils.GetEnv()

func main() {
	// Checking if no commands are passed in and if the binary is simply being called...
	if len(os.Args) < 2 {
		var out bytes.Buffer
		cmd := exec.Command(fmt.Sprintf("%s/cmdbin/help", ENV["BURAQROOT"]))
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(out.String())
		os.Exit(0)
	}
}
