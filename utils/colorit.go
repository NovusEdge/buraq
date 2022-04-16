package utils

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
)

// Colors for console use
const (

	// Normal text colors
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorGrey   = "\033[1;30m"

	// Bold text colors
	BoldColorRed    = "\033[1;31m"
	BoldColorGreen  = "\033[1;32m"
	BoldColorYellow = "\033[1;33m"
	BoldColorBlue   = "\033[1;34m"
	BoldColorPurple = "\033[1;35m"
	BoldColorCyan   = "\033[1;36m"
	BoldColorWhite  = "\033[1;37m"
	BoldColorGrey   = "\033[1;30m"

	// Dim text colors
	DimColorRed    = "\033[2;31m"
	DimColorGreen  = "\033[2;32m"
	DimColorYellow = "\033[2;33m"
	DimColorBlue   = "\033[2;34m"
	DimColorPurple = "\033[2;35m"
	DimColorCyan   = "\033[2;36m"
	DimColorWhite  = "\033[2;37m"
	DimColorGrey   = "\033[2;30m"

	// Italic text colors
	ItalicColorRed    = "\033[3;31m"
	ItalicColorGreen  = "\033[3;32m"
	ItalicColorYellow = "\033[3;33m"
	ItalicColorBlue   = "\033[3;34m"
	ItalicColorPurple = "\033[3;35m"
	ItalicColorCyan   = "\033[3;36m"
	ItalicColorWhite  = "\033[3;37m"
	ItalicColorGrey   = "\033[3;30m"

	// Underlined text colors
	UnderlineColorRed    = "\033[4;31m"
	UnderlineColorGreen  = "\033[4;32m"
	UnderlineColorYellow = "\033[4;33m"
	UnderlineColorBlue   = "\033[4;34m"
	UnderlineColorPurple = "\033[4;35m"
	UnderlineColorCyan   = "\033[4;36m"
	UnderlineColorWhite  = "\033[4;37m"
	UnderlineColorGrey   = "\033[4;30m"

	// Blink text colors
	BlinkColorRed    = "\033[5;31m"
	BlinkColorGreen  = "\033[5;32m"
	BlinkColorYellow = "\033[5;33m"
	BlinkColorBlue   = "\033[5;34m"
	BlinkColorPurple = "\033[5;35m"
	BlinkColorCyan   = "\033[5;36m"
	BlinkColorWhite  = "\033[5;37m"
	BlinkColorGrey   = "\033[5;30m"

	// Negative text colors
	NegativeColorRed    = "\033[7;31m"
	NegativeColorGreen  = "\033[7;32m"
	NegativeColorYellow = "\033[7;33m"
	NegativeColorBlue   = "\033[7;34m"
	NegativeColorPurple = "\033[7;35m"
	NegativeColorCyan   = "\033[7;36m"
	NegativeColorWhite  = "\033[7;37m"
	NegativeColorGrey   = "\033[7;30m"

	// Crossed text colors
	CrossedColorRed    = "\033[9;31m"
	CrossedColorGreen  = "\033[9;32m"
	CrossedColorYellow = "\033[9;33m"
	CrossedColorBlue   = "\033[9;34m"
	CrossedColorPurple = "\033[9;35m"
	CrossedColorCyan   = "\033[9;36m"
	CrossedColorWhite  = "\033[9;37m"
	CrossedColorGrey   = "\033[9;30m"

	// Simple Modifiers
	Bold      = "\033[1m"
	Dim       = "\033[2m"
	Italic    = "\033[3m"
	Underline = "\033[4m"
	Blink     = "\033[5m"
	Negative  = "\033[7m"
	Crossed   = "\033[9m"
)

// ColorIt applies color to a string.
func ColorIt(c string, s string) string {
	return fmt.Sprintf("%s%s%s", c, s, ColorReset)
}

// ApplyEffect applies a particular effect to the string.
// This is the same as ColorIt, but is defined for making the code
// more organized and managable.
func ApplyEffect(e string, s string) string {
	return fmt.Sprintf("%s%s%s", e, s, ColorReset)
}
