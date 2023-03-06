package buraqlib

import (
	"fmt"
)

func PrintError(message string) {
	fmt.Println(ColorIt(ColorRed, "[-]: "+message))
}

func PrintWarning(message string) {
	fmt.Println(ColorIt(ColorYellow, "[!]: "+message))
}

func PrintInfo(message string) {
	fmt.Println(ColorIt(ColorGrey, "[*]: "+message))
}

func PrintSuccess(message string) {
	fmt.Println(ColorIt(ColorGreen, "[+]: "+message))
}
