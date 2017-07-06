package mydb

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	Models "github.com/tknott95/MasterGo/Models"
)

var Store = newDB()

func newDB() *Models.SQLStore {
	db, err := sql.Open("mysql", "tknott95:Welcome1!@tcp(admininstance.cfchdss74ohb.us-west-1.rds.amazonaws.com:3306)/adminaws?charset=utf8")
	if err != nil {
		println("🔒 Connection to AWS database established.\n")
	}

	return &Models.SQLStore{
		DB: db,
	}
}
