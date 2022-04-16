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
)

func main() {
	hd := utils.GetHomeDirectory()
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/.buraq/env", hd))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
