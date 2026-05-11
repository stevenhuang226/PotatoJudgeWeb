package db

import "database/sql"

type DB struct {
	Conn *sql.DB
}

func (db *DB) init() {
}
