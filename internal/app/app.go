package app

import (
	"database/sql"
	"errors"
	"pjweb/internal/config"
)

type App struct {
	Debug    bool
	Problem  Problem
	Submit   Submit
	Static   Static
	Database Database
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
	BasePath           string
	SocketPath         string
	MaxQueueSize       int32
	MaxConcurrentJudge int32
	PJCompilerPrefix   string
	PJDetailName       string
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

	app.Debug = cfg.Debug

	/* problem */
	if cfg.Problem.BasePath == "" {
		return nil, errors.New("no problem path")
	}
	app.Problem.BasePath = cfg.Problem.BasePath
	if cfg.Problem.InCasePrefix != "" {
		app.Problem.InCasePrefix = cfg.Problem.InCasePrefix
	} else {
		app.Problem.InCasePrefix = "input"
	}
	if cfg.Problem.InCaseSuffix != "" {
		app.Problem.InCaseSuffix = cfg.Problem.InCaseSuffix
	} else {
		app.Problem.InCaseSuffix = ".bin"
	}
	if cfg.Problem.OutCasePrefix != "" {
		app.Problem.OutCasePrefix = cfg.Problem.OutCasePrefix
	} else {
		app.Problem.OutCasePreifx = "output"
	}
	if cfg.Problem.OutCaseSuffix != "" {
		app.Problem.OutCaseSuffix = cfg.Problem.OutCaseSuffix
	} else {
		app.Problem.OutCaseSuffix = ".bin"
	}

	/* submit */
	if cfg.Submit.BasePath == "" || cfg.Submit.SocketPath == "" {
		return nil, errors.New("no submit/socket path")
	}
	app.Submit.BasePath = cfg.Submit.BasePath
	app.Submit.SocketPath = cfg.Submit.SocketPath
	app.Submit.PJCompilerPrefix = "compiler_type="
	if cfg.Submit.MaxConcurrentJudge > 0 {
		app.Submit.MaxQueueSize = cfg.Submit.MaxConcurrentJudge
	} else {
		app.Submit.MaxQueueSize = 128
	}
	if cfg.Submit.MaxConcurrentJudge > 0 {
		app.Submit.MaxConcurrentJudge = cfg.Submit.MaxConcurrentJudge
	} else {
		app.Submit.MaxConcurrentJudge = 8
	}

	/* static */
	if cfg.Static.BasePath == "" {
		return nil, errros.New("no static path")
	}
	app.Static.BasePath = cfg.Static.BasePath

	return &app
}
