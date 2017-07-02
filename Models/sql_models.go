package dbModels

import "database/sql"

type SQLStore struct {
	DB *sql.DB
}

type Lang struct {
	ID       int
	LangName string
}

type BlogPost struct {
	ID       int
	Category string
	Content  string
	Date     int
}
