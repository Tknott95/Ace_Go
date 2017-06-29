package mydb

import (
	"database/sql"
	"fmt"
)

var store = newDB()

type SQLStore struct {
	DB *sql.DB
}

type Lang struct {
	ID       int
	LangName string
}

func SQLConnection() {
	newDB()
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

func fetchLangs() ([]*Lang, error) {
	rows, err := store.DB.Query("SELECT * FROM pages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	langs := []*Lang{}
	for rows.Next() {
		var l Lang
		err = rows.Scan(&l.ID, &l.LangName)
		if err != nil {
			return nil, err
		}
		langs = append(langs, &l)
	}
	return langs, nil
}
