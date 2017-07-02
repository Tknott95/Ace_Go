package main

import (
	globals "github.com/tknott95/Inspired/Globals"
	srvCtrl "github.com/tknott95/MasterGo/Controllers"
)

func main() {
	/* Intro in Term */
	println("\n || Trevor Knott Admin || \n ______________________ \n")

	/* On Port ? */
	println("\n🚀 Server Running on Port: " + globals.PortNumber + "\n")

	/* Begin Server w/ Routes */
	srvCtrl.InitServer()
}
