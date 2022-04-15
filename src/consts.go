package gydeon

/*
// Author: Aliasgar Khimani (NovusEdge)
// Project: github.com/NovusEdge/gydeon
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

// Version for the gydeon CLI
const Version string = "v1.0.0"

// HelpScreen lists all command line flags and usage for gydeon.
const HelpScreen string = `
Gydeon:
    A command-line utility for bruteforcing ssh

Usage:
    gydeon [command] [options] target

Commands:
    check               perform simple enumeration for ssh services on target.

Options:
    -h --help           prints this menu
    -v --version        prints the version of the program
    -q --quiet          enable output suppression
`
