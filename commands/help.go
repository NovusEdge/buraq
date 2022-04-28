package main

import (
	"fmt"

	src "github.com/NovusEdge/buraq/src"
)

func main() {
	fmt.Println("buraq", src.Version, "(c) 2022 by Aliasgar Khimani (NovusEdge) - Please do not use in military or secret service organizations, or for illegal purposes (this is non-binding, these *** ignore laws and ethics anyway).")
	fmt.Println(src.HelpScreen)
}
