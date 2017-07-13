#!/bin/bash

# Assuming you already got project
cd $GOPATH/src/github.com/tknott95/Ace_Go/

# Add timestamp to name end @TODO
rm -f Ace_Go
go build

go run application.go
