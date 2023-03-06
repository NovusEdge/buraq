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
	"os"

	buraqlib "github.com/NovusEdge/buraq/buraqlib"
	src "github.com/NovusEdge/buraq/src"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("buraq", src.Version, "(c) 2022 by Aliasgar Khimani (NovusEdge) - Please do not use in military or secret service organizations, or for illegal purposes (this is non-binding, these *** ignore laws and ethics anyway).")
		fmt.Println(src.HelpScreen)
	} else {
		cmd := os.Args[1]
		switch cmd {
		case "help":
			fmt.Println(src.CommandHelpHelp)
		case "env":
			fmt.Println(src.CommandEnvHelp)
		case "attack":
			fmt.Println(src.CommandAttackHelp)
		case "version":
			fmt.Println(src.CommandVersionHelp)
		case "clean":
			fmt.Println(src.CommandCleanHelp)
		case "repair":
			fmt.Println(src.CommandRepairHelp)
		default:
			fmt.Println(buraqlib.ColorIt(buraqlib.ColorRed, fmt.Sprintf("[E]: Invalid Command: '%s'\nPlease use 'buarq help' for usage", cmd)))
		}
	}
}
