package db

import (
	"context"
	"database/sql"
	"errors"
	"pjweb/internal/app"
	"pjweb/internal/config"
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
	connMaxLifetime := time.hour

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

	return &conn, nil
}

func (app *app.App) InitDB(cfg *config.DatabaseConfig) error {
	if cfg == nil {
		return errors.New("nil DatabaseConfig")
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
		return err
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
		return err
	}

	app.Database.Conn = conn
	return nil
}
