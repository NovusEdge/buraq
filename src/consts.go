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
 ██████╗██╗   ██╗██████╗ ███████╗ ██████╗ ███╗   ██╗
██╔════╝╚██╗ ██╔╝██╔══██╗██╔════╝██╔═══██╗████╗  ██║
██║  ███╗╚████╔╝ ██║  ██║█████╗  ██║   ██║██╔██╗ ██║
██║   ██║ ╚██╔╝  ██║  ██║██╔══╝  ██║   ██║██║╚██╗██║
╚██████╔╝  ██║   ██████╔╝███████╗╚██████╔╝██║ ╚████║
 ╚═════╝   ╚═╝   ╚═════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝`

// Version for the buraq CLI
const Version string = "v1.0.0"

// HelpScreen lists all command line flags and usage for buraq.
const HelpScreen string = `
Buraq:
    A command-line utility for bruteforcing ssh

USAGE:
    buraq [command] [options] target

COMMANDS:
    check               perform simple enumeration for ssh services on target.

OPTIONS:
    -h --help           prints this menu
    -v --version        prints the version of the program
    -q --quiet          enable output suppression
`
