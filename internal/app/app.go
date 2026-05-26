package app

import (
	"database/sql"
	"errors"
	"pjweb/internal/api"
	"pjweb/internal/config"
	"pjweb/internal/db"
	"pjweb/internal/repository"
	"pjweb/internal/service"
)

type App struct {
	Debug    Debug
	Problem  Problem
	Submit   Submit
	Static   Static
	Database Database
	Server   Server

	ProblemSvc *service.ProblemService
	SubmitSvc  *service.SubmitService
	StaticSvc  *service.StaticService

	ProblemRepo  *repository.ProblemRepo
	SubmitRepo   *repository.SubmitRepo
	StaticRepo   *repository.StaticRepo
	DatabaseRepo *repository.DatabaseRepo

	Router *api.Router
}

type Debug struct {
	IsOn  bool
	UseDB bool
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

type Server struct {
	Port uint16
	Host string
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

	/* debug section */
	if cfg.Debug.IsDebug && !cfg.Debug.UseDB {
		return &app, nil
	}
	/* debug section end */

	app.Database.Conn, err = db.SetUpDB(
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
