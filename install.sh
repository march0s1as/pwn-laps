#!/usr/bin/env bash

# cores
normal=$'\e[0m'  
C=$(printf '\033')                                                 
green="${C}[1;32m"
yellow="${C}[1;33m"
RED="${C}[1;31m"
# fim das cores

instalacao(){
	echo "${green}[OK] ${normal}starting the installation."
	sleep 2
	echo "${yellow}[!!] ${normal}looking for the go path."

	if [ -x "$(command -v go)" ]; then
		echo "${green}[OK] ${normal}SUCCESS."
		go get -v "github.com/fatih/color"
		go get -v "gopkg.in/ldap.v2"
		go build laps.go
		echo "${green}[OK] ${normal}to use this tool, type ${yellow}./laps${normal} and have fun! =]"

	else
		echo "${RED}plz install golang lol"
	fi
}

instalacao