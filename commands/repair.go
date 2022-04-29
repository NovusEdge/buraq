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
	"fmt"
	utils "github.com/NovusEdge/buraq/utils"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	env := utils.GetEnv()

	// Cleaning binaries...
	cleaningErr := exec.Command(fmt.Sprintf("%s/clean", env["BURAQCMDBIN"])).Run()
	if cleaningErr != nil {
		log.Fatal(cleaningErr)
	}

	// Rebuilding binaries...
	os.Chdir(env["BURAQROOT"])

	mainErr := exec.Command("go", "build", "-o", fmt.Sprintf("%s/bin/", env["BURAQROOT"]), fmt.Sprintf("%s/buraq.go", env["BURAQROOT"])).Run()
	if mainErr != nil {
		log.Fatal(mainErr)
	}

	cmds, err := ioutil.ReadDir(fmt.Sprintf("%s/commands", env["BURAQROOT"]))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range cmds {
		err = exec.Command("go", "build", "-o", fmt.Sprintf("%s/cmdbin/", env["BURAQROOT"]), fmt.Sprintf("%s/commands/%s", env["BURAQROOT"], file.Name())).Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
