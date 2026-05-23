package repository

import (
	"database/sql"
	"pjweb/internal/app"
)

type DatabaseRepo struct {
	Conn *sql.DB
}

func NewDatabase(db *app.Database) (*DatabaseRepo, error) {
	var repo DatabaseRepo

	repo.Conn = db.Conn

	return &repo, nil
}
