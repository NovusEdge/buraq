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
		"clean",
		"attack",
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

    attack              Perform a bruteforce attack on specified target using
                        a user/userlist and a password list.
                        If no user is specified, "root" is used.
                        If no password list is used, "~/.buraq/passlist.txt" is used.

    clean               Clear all binaries for the CLI.
    repair              Clear and rebuild all binaries for the CLI.`

// The help/manuals for each command.
const (
	CommandAttackHelp = `USAGE:
	buraq attack [options] [target]

OPTIONS:
    -port
        Specifies the port to use for the attack. (Default: 22)

    -proto
        Specifies the protocol to use for the attack (tcp/udp) (Default: tcp)

    -user
        Specifies the user/login-name for the attack. (Default: root)

    -userlist
        Specifies a user-list to use for the attack.
        If this is specified along with the -u option then the value from that will be appended to the user-list.

    -passlist
        Specifies a password-list to use for the attack. (Default: ~/.buraq/passlist.txt)
        NOTE: The specified password list is the same as the one found on seclists:
        https://github.com/danielmiessler/SecLists/blob/master/Passwords/xato-net-10-million-passwords-100000.txt

    -timeout
        Specifies the timeout between each attack attempt in milliseconds. (Default: 500ms)

    -threads
        Specify the number of threads to use for the attack. (Default: 16)`

	CommandHelpHelp = `USAGE:
    buraq help [command]

The help, if used on it's own (e.g. 'buraq help') displays usage for the CLI.
When supplied with the name of another command, the help command displays information about the said command.

Example:
    buraq help attack

This will display the usage information about the 'attack' command.`

	CommandVersionHelp = `USAGE:
    buraq version

The 'version' command reports the version of buraq in use.`
	CommandEnvHelp = `
USAGE:
    buraq env

The 'env' command prints the environment variables used by buraq.
These variables are stored in ~/.buraq/env (given that the setup file is run)`

	CommandRepairHelp = `USAGE:
    buraq repair

The 'repair' command rebuilds all binaries that are used by the CLI.
Note that it does not fetch and/or update the golang dependencies,
to do that you must run the setup script.`

	CommandCleanHelp = `USAGE:
    buraq clean

The 'clean' command deletes/removes all binaries compiled and used by buraq.`
)
