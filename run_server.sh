#!/bin/bash

# Assuming you already got project
cd $GOPATH/src/github.com/tknott95/MasterGo/

go run fmt */**/*.go
go run main.go
