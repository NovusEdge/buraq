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
)

func main() {
	env := utils.GetEnv()
	bin, err := ioutil.ReadDir(env["GYDEONBIN"])
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range bin {
		os.RemoveAll(path.Join([]string{env["GYDEONBIN"], file.Name()}...))
	}
}
