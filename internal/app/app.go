package app

import (
	"database/sql"
	"errors"
	"pjweb/internal/app"
	"pjweb/internal/config"
	"pjweb/internal/db"
	"pjweb/internal/repository"
	"pjweb/internal/service"
)

type App struct {
	Debug    bool
	Problem  Problem
	Submit   Submit
	Static   Static
	Database Database

	ProblemSvc *service.ProblemService
	SubmitSvc  *service.SubmitService
	StaticSvc  *service.StaticService

	ProblemRepo  *repository.ProblemRepo
	SubmitRepo   *repository.SubmitRepo
	StaticRepo   *repository.StaticRepo
	DatabaseRepo *repository.DatabaseRepo
}

type Problem struct {
	BasePath      string
	Explanation   string
	InCasePrefix  string
	InCaseSuffix  string
	OutCasePrefix string
	OutCaseSuffix string
}

type Submit struct {
	BasePath             string
	SocketPath           string
	MaxQueueSize         int32
	MaxConcurrentJudge   int32
	PJCompilerTypePrefix string
	PJDetailName         string
}

type Database struct {
	Conn *sql.DB
}

type Static struct {
	BasePath string
}

func New(cfg *config.Config) (*App, error) {
	if cfg == nil {
		return nil, errors.New("no cfg")
	}

	var app App

	err := app.SetFromConfig(cfg)
	if err != nil {
		return nil, err
	}

	app.Database.Conn, err := db.SetUpDB(
		cfg.Database.DSN,
		cfg.Database.MaxOpenConns,
		cfg.Database.MaxIdleConns,
		cfg.Database.ConnMaxLifetime,
	)

	if err != nil {
		return nil, err
	}

	return &app, nil
}
