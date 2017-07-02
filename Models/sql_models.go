package dbModels

import "database/sql"

type SQLStore struct {
	DB *sql.DB
}

type Lang struct {
	ID       int
	LangName string
}
