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
	"flag"
	"os"
	"log"
	src "github.com/NovusEdge/buraq/src"
	utils "github.com/NovusEdge/buraq/utils"
	"os/exec"
    "io/ioutil"
)

// ENV is the environment variables for the program
var ENV = utils.GetEnv()

func main() {
	var quiet bool
	var port uint
	var target string
    flag.Usage = func() {
        fmt.Println("buraq", src.Version, "(c) 2022 by Aliasgar Khimani (NovusEdge) - Please do not use in military or secret service organizations, or for illegal purposes (this is non-binding, these *** ignore laws and ethics anyway).")
        fmt.Println(src.HelpScreen)
    }

    if (len(os.Args) == 1) {
        PrintHelpScreen()
        os.Exit(0)
    }

    command := os.Args[1]

    switch command {
    case "help":
        PrintHelpScreen()
        os.Exit(0)

    case "version":
        fmt.Println("buraq", src.Version)
        os.Exit(0)

    case "env":
        PrintEnv()
        os.Exit(0)

    case "repair":
        RepairBinaries()
        os.Exit(0)
    }


    flag.BoolVar(&quiet, "q", false, "Enable output suppression for commands")
    flag.BoolVar(&quiet, "quiet", false, "Enable output suppression for commands")

    flag.UintVar(&port, "p", 22, "Specifies the port on which the ssh service is running")
    flag.UintVar(&port, "port", 22, "Specifies the port on which the ssh service is running")

    // implement checking for valid target string...

    target = os.Args[len(os.Args)-1]
    flag.Parse()

    switch command {
    case "check":
        // src.CheckTarget()
        fmt.Println("checking :)")
        os.Exit(0)
    }
    fmt.Println(target)
    fmt.Println(command, quiet, port)
    exec.Command("whoami").Run()
}

// PrintBanner prints the ascii banner art for buraq.
func PrintBanner() {
	fmt.Println(src.BannerArt)
}

// PrintHelpScreen prints the help screen for buraq.
func PrintHelpScreen() {
	fmt.Println("buraq", src.Version, "(c) 2022 by Aliasgar Khimani (NovusEdge) - Please do not use in military or secret service organizations, or for illegal purposes (this is non-binding, these *** ignore laws and ethics anyway).")
	fmt.Println(src.HelpScreen)
}

// PrintEnv prints environment for the program
func PrintEnv() {
    hd := utils.GetHomeDirectory()
    content, err := ioutil.ReadFile(fmt.Sprintf("%s/.buraq/env", hd))
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(content))
}

// RepairBinaries evokes the repair binary
func RepairBinaries() {
    err := exec.Command(fmt.Sprintf("%s/repair", ENV["BURAQBIN"])).Run()
    if err != nil {
        log.Fatal(err)
    }
}
