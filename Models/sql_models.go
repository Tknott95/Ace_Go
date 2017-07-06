package Models

import "database/sql"

type SQLStore struct {
	DB *sql.DB
}

