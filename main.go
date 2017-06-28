package main

import (
	globals "github.com/tknott95/Inspired/Globals"
	srvCtrl "github.com/tknott95/MasterGo/Controllers"
	mydb "github.com/tknott95/MasterGo/Controllers/DB_Ctrl"
)

func main() {
	/* Intro in Term */
	println("\n || Trevor Knott Admin || \n ______________________ \n")

	/* On Port ? */
	println("\n🚀 Server Running on Port: " + globals.PortNumber + "\n")

	/* Est a mySql Connection before serving pages */
	mydb.SQLConnection()

	/* Begin Server w/ Routes */
	srvCtrl.InitServer()
}
