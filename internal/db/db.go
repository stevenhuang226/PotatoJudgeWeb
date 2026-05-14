package db

import (
	"context"
	"database/sql"
	"errors"
	"pjweb/internal/config"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	Conn *sql.DB
}

func New(cfg *config.DatabaseConfig) (*Database, error) {
	if cfg == nil {
		return nil, errors.New("nil DatabaseConfig")
	}

	const (
		DefMaxOpenConns    int           = 16
		DefMaxIdleConns    int           = 8
		DefConnMaxLifetime time.Duration = time.Hour
	)
	MaxOpenConns := cfg.MaxOpenConns
	if MaxOpenConns <= 0 {
		MaxOpenConns = DefMaxOpenConns
	}
	MaxIdleConns := cfg.MaxIdleConns
	if MaxIdleConns <= 0 {
		MaxIdleConns = DefMaxIdleConns
	}
	ConnMaxLifetime := cfg.MaxConnLifetime
	if ConnMaxLifetime <= 0 {
		ConnMaxLifetime = DefConnMaxLifetime
	}

	conn, err := sql.Open(
		"pgx",
		cfg.DSN,
	)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(MaxOpenConns)
	conn.SetMaxIdleConns(MaxIdleConns)
	conn.SetConnMaxLifetime(ConnMaxLifetime)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := conn.PingContext(ctx); err != nil {
		conn.Close()
		return nil, err
	}

	return &Database{
		Conn: conn,
	}, nil
}
