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
	"log"
	"os"
	"os/exec"
)

func main() {
	env := utils.GetEnv()

	// Cleaning binaries...
	err := exec.Command(fmt.Sprintf("%s/clean", env["BURAQBIN"])).Run()
	if err != nil {
		log.Fatal(err)
	}

	// Rebuilding binaries...
	os.Chdir(env["BURAQROOT"])

	err = exec.Command("go", "build", "-o", fmt.Sprintf("%s/bin/", env["BURAQROOT"]), fmt.Sprintf("%s/buraq.go", env["BURAQROOT"])).Run()
	if err != nil {
		log.Fatal(err)
	}

	err = exec.Command("go", "build", "-o", fmt.Sprintf("%s/bin/", env["BURAQROOT"]), fmt.Sprintf("%s/tools/clean.go", env["BURAQROOT"])).Run()
	if err != nil {
		log.Fatal(err)
	}

	err = exec.Command("go", "build", "-o", fmt.Sprintf("%s/bin/", env["BURAQROOT"]), fmt.Sprintf("%s/tools/repair.go", env["BURAQROOT"])).Run()
	if err != nil {
		log.Fatal(err)
	}
}
