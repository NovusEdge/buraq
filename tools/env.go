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
	"fmt"
	utils "github.com/NovusEdge/gydeon/utils"
	"io/ioutil"
	"log"
)

func main() {
	hd := utils.GetHomeDirectory()
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/.gydeon/env", hd))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
