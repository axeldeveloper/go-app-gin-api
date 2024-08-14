#!/bin/bash

# This script displays help information for the Makefile.
# Usage: ./help.sh Makefile

# Set colors for output
normal=$'\e[0m'                           # (works better sometimes)
bold=$(tput bold)                         # make colors bold/bright
red="$bold$(tput setaf 1)"                # bright red text
green=$(tput setaf 2)                     # dim green text
fawn=$(tput setaf 3); beige="$fawn"       # dark yellow text
yellow="$bold$fawn"                       # bright yellow text
darkblue=$(tput setaf 4)                  # dim blue text
blue="$bold$darkblue"                     # bright blue text
purple=$(tput setaf 5); magenta="$purple" # magenta text
pink="$bold$purple"                       # bright magenta text
darkcyan=$(tput setaf 6)                  # dim cyan text
cyan="$bold$darkcyan"                     # bright cyan text
gray=$(tput setaf 7)                      # dim white text
darkgray="$bold"$(tput setaf 0)           # bold black = dark gray text
white="$bold$gray"                        # bright white text

PROMPT_COMMAND='[ $? = 0 ] && PS1=${PS1[1]} || PS1=${PS1[2]}'

# Main function to display help information
help() {
    # Display usage information
    echo "Usage:"
    printf "  make [target] [variables]\n\n"
    echo "${red}hello ${yellow}this is ${green}coloured${normal}"
    
    # Display targets information
    #_help_targets "$1"
    
    # Display variables information
    # _help_variables "$1"
    
    # Display examples
    #_help_examples
}

_run() {
    
    echo "${yellow}Rum Project(s): ${yellow} go ${red}hello"
    go run server.go 
}


set_path() {
     echo "${yellow}Set go in path(s): ${yellow} go ${red}hello"
    # export PATH=$PATH:$(go env GOPATH)/bin
    export PATH=$(go env GOPATH)/bin:$PATH
}

_init(){
    echo "${yellow} initialized: ${yellow} swag ${red} init"
    swag init -g server.go
}

# Call main function
#help "$1"
"$@"

# go get -u github.com/gin-gonic/gin
# go get -u github.com/swaggo/swag/cmd/swag@latest
# go get -u github.com/swaggo/gin-swagger
# go get -u github.com/swaggo/files