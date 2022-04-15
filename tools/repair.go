package main

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
	utils "github.com/NovusEdge/gydeon/utils"
	"io/ioutil"
	"log"
	"os"
	"path"
    "os/exec"
    "fmt"
)

func main() {
	env := utils.GetEnv()

    // Cleaning binaries...
    bin, err := ioutil.ReadDir(env["GYDEONBIN"])
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range bin {
		os.RemoveAll(path.Join([]string{env["GYDEONBIN"], file.Name()}...))
	}

    // Rebuilding binaries...
    os.Chdir(env["GYDEONROOT"])

    err = exec.Command("go", "build", "-o", fmt.Sprintf("%s/bin/", env["GYDEONROOT"]), fmt.Sprintf("%s/gydeon.go", env["GYDEONROOT"])).Run()
    if err != nil {
        log.Fatal(err)
    }

    err = exec.Command("go", "build", "-o", fmt.Sprintf("%s/bin/", env["GYDEONROOT"]), fmt.Sprintf("%s/tools/clean.go", env["GYDEONROOT"])).Run()
    if err != nil {
        log.Fatal(err)
    }

    err = exec.Command("go", "build", "-o", fmt.Sprintf("%s/bin/", env["GYDEONROOT"]), fmt.Sprintf("%s/tools/repair.go", env["GYDEONROOT"])).Run()
    if err != nil {
        log.Fatal(err)
    }
}
