package mydb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var store = newDB()

type SQLStore struct {
	DB *sql.DB
}

type Lang struct {
	ID       int
	LangName string
}

func newDB() *SQLStore {
	db, err := sql.Open("mysql", "tknott95:Welcome1!@tcp(admininstance.cfchdss74ohb.us-west-1.rds.amazonaws.com:3306)/adminaws?charset=utf8")
	if err != nil {
		fmt.Println("🔒 Connection to AWS database established.\n")
	}

	return &SQLStore{
		DB: db,
	}
}

func FetchLangs() []*Lang {
	rows, err := store.DB.Query("SELECT * FROM pc_langs;")
	if err != nil {
		return nil
	}
	defer rows.Close()

	langs := []*Lang{}
	for rows.Next() {
		var l Lang
		err = rows.Scan(&l.ID, &l.LangName)
		if err != nil {
			return nil
		}
		langs = append(langs, &l)
	}
	return langs
}
