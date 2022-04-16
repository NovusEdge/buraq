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
	utils "github.com/NovusEdge/buraq/utils"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	env := utils.GetEnv()
	bin, err := ioutil.ReadDir(env["BURAQBIN"])
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range bin {
		os.RemoveAll(path.Join([]string{env["BURAQBIN"], file.Name()}...))
	}
}
