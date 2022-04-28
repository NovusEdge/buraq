package buraq

/*
// Author: Aliasgar Khimani (NovusEdge)
// Project: github.com/NovusEdge/buraq
//
// Copyright: GNU General Public License v3.0
// See the LICENSE file for more info.
//
// All Rights Reserved
*/

// BannerArt for the CLI
const BannerArt string = `
██████  ██    ██ ██████   █████   ██████
██   ██ ██    ██ ██   ██ ██   ██ ██    ██
██████  ██    ██ ██████  ███████ ██    ██
██   ██ ██    ██ ██   ██ ██   ██ ██ ▄▄ ██
██████   ██████  ██   ██ ██   ██  ██████
                                     ▀▀`

// Version for the buraq CLI
const Version string = "v1.0.0"

// ValidCommands that can be provided to the CLI.
func ValidCommands() []string {
	return []string{
		"env",
		"help",
		"check",
		"clean",
		"repair",
		"version",
	}
}

// HelpScreen lists all command line flags and usage for buraq.
const HelpScreen string = `
Buraq:
    A command-line utility for bruteforcing ssh

USAGE:
    buraq [command] [options] target

COMMANDS:
    help                Prints this menu if no further option is provided. This
                        can be used in conjunction with other command names to
                        print a more comprehensive guide for usage of said command.
                        For example:
                            buraq help check
                        This will print out information about the check command.

    version             Prints the version of buraq in use.
    env                 Prints the environment variables used by the CLI.
                        These are stored (by default) in ~/.buraq

    check               Perform simple enumeration for ssh services on target.
    clean               Clear all binaries for the CLI.
    repair              Clear and rebuild all binaries for the CLI.`
