package main

import (
	"fmt"
	"io/ioutil"
	"log"

	utils "github.com/NovusEdge/buraq/utils"
)

func main() {
	hd := utils.GetHomeDirectory()
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/.buraq/env", hd))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
