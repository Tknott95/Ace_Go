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
	Title    string
	Image    string
	Category string
	Content  string
	Author   string
	Date     int
}
