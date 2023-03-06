#!/usr/bin/env bash

: '
    Author: Aliasgar Khimani (NovusEdge)
    Project: github.com/NovusEdge/buraq
    Copyright: GNU General Public License v3.0
    See the LICENSE file for more info.
    All Rights Reserved
'

VERSION="v1.0.0"

printf "\033[1;36m=================LICENSE NOTICE=================\n";
echo "\
    Author: Aliasgar Khimani (NovusEdge)
    Project: github.com/NovusEdge/buraq
    Version: $VERSION
    Copyright: GNU General Public License v3.0
    See the LICENSE file for more info.
    All Rights Reserved
=================================================
"; printf "\033[0m"


PROJECT_DIR=$(dirname $(realpath $0))

if [ ! -d "$PROJECT_DIR/bin" ]; then
    printf "\033[1;36m[I]: Making the binary directory\033[0m\n\tBinary directory: $PROJECT_DIR/bin/\n\n";
    mkdir -p $PROJECT_DIR/bin;
fi
if [ ! -d "$PROJECT_DIR/cmdbin" ]; then
    printf "\033[1;36m[I]: Making the command-binary directory\033[0m\n\tBinary directory: $PROJECT_DIR/bin/cmdbin\n\n";
    mkdir -p $PROJECT_DIR/cmdbin;
fi

if ! type "go" > /dev/null; then
    printf "\033[1;31m[E]: Golang not detected!\n";
    printf "     Please install it from: https://go.dev/ \033[0m\n\n";
    exit 1;
else
    printf "\033[1;32m[+]: Found golang installation at: \n\t\033[35m$(which go)\033[0m\n\n";
    GO=$(which go)
fi


## Getting golang dependencies...
oldcwd=$(pwd)
printf "\033[1;36m[I]: Getting golang dependencies...\033[0m\n";

cd $PROJECT_DIR
go get -x
go get -x -d ./src/
go get -x -d ./utils/
go get -x -d ./commands/
cd $oldcwd

printf "\033[1;32m[+]: Done!\033[0m\n\n";

## Building binaries...
printf "\033[1;36m[I]: Building binaries...\033[0m\n";

# Building main binary...
$GO build -o $PROJECT_DIR/bin/ $PROJECT_DIR/buraq.go

# Building command binaries...
$GO build -o $PROJECT_DIR/cmdbin/ $PROJECT_DIR/commands/env.go
$GO build -o $PROJECT_DIR/cmdbin/ $PROJECT_DIR/commands/help.go
$GO build -o $PROJECT_DIR/cmdbin/ $PROJECT_DIR/commands/clean.go
$GO build -o $PROJECT_DIR/cmdbin/ $PROJECT_DIR/commands/attack.go
$GO build -o $PROJECT_DIR/cmdbin/ $PROJECT_DIR/commands/repair.go
$GO build -o $PROJECT_DIR/cmdbin/ $PROJECT_DIR/commands/version.go

printf "\033[1;32m[+]: Done!\033[0m\n\n";


## Making ~/.buraq if it does not exist.
if [ ! -d $HOME/.buraq ]; then
    printf "\033[1;36m[I]: Making .buraq in $HOME for storing config files and logs...\033[0m\n";
    mkdir $HOME/.buraq
    touch $HOME/.buraq/env
    printf "\033[1;32m[+]: Done!\033[0m\n\n";
fi

## Setting up ~/.buraq/env
printf "\033[1;36m[I]: Setting up environment file...\033[0m\n";

echo BURAQROOT="$PROJECT_DIR"             > $HOME/.buraq/env
echo BURAQVER="$VERSION"                 >> $HOME/.buraq/env
echo BURAQENV="$HOME/.buraq/env"         >> $HOME/.buraq/env
echo BURAQBIN="$PROJECT_DIR/bin"         >> $HOME/.buraq/env
echo BURAQCMDBIN="$PROJECT_DIR/cmdbin"   >> $HOME/.buraq/env
echo GOVERSION="$($GO version)"          >> $HOME/.buraq/env

printf "\033[1;32m[+]: Done!\033[0m\n\n";


cp $PROJECT_DIR/default_passlist.txt $HOME/.buraq/
cp $PROJECT_DIR/default_userlist.txt $HOME/.buraq/
