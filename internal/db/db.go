package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func SetUpDB(
	dsn string,
	_maxOpenConns int,
	_maxIdleConns int,
	_connMaxLifetime time.Duration,
) (*sql.DB, error) {
	if dsn == "" {
		return nil, errors.New("no dsn")
	}

	maxOpenConns := 16
	maxIdleConns := 8
	connMaxLifetime := time.Hour

	if _maxOpenConns > 0 {
		maxOpenConns = _maxOpenConns
	}
	if _maxIdleConns > 0 {
		maxIdleConns = _maxIdleConns
	}
	if _connMaxLifetime > 0 {
		connMaxLifetime = _connMaxLifetime
	}

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(maxOpenConns)
	conn.SetMaxIdleConns(maxIdleConns)
	conn.SetConnMaxIdleTime(connMaxLifetime)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := conn.PingContext(ctx); err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}
