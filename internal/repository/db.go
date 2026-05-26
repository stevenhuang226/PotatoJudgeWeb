package repository

import (
	"database/sql"
)

type DatabaseRepo struct {
	Conn *sql.DB
}

func NewDatabaseRepo(conn *sql.DB) (*DatabaseRepo, error) {
	return &DatabaseRepo{
		Conn: conn,
	}, nil
}
