#!/usr/bin/env bash

: '
    Author: Aliasgar Khimani (NovusEdge)
    Project: github.com/ARaChn3/buraq
    Copyright: GNU General Public License v3.0
    See the LICENSE file for more info.
    All Rights Reserved
'

PROJECT_DIR=$(dirname $(realpath $0))

printf "\033[1;36m[I]: Removing $HOME/.buraq...\033[0m\n";

rm -rf $HOME/.buraq

printf "\033[1;32m[+]: Done!\033[0m\n\n";



printf "\033[1;36m[I]: Removing the project directory and files...\n";

rm -rf $PROJECT_DIR

printf "\033[1;32m[+]: Done!\033[0m\n\n";
