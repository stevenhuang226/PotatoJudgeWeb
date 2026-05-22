package service

import (
	"database/sql"
	"errors"
	"pjweb/internal/config"
)

type Service struct {
	Debug    bool
	Problem  ProblemSvc
	Submit   SubmitSvc
	Database DatabaseSvc
}

type ProblemSvc struct {
	BasePath      string
	Explanation   string
	InCasePrefix  string
	InCaseSuffix  string
	OutCasePrefix string
	OutCaseSuffix string
}

type SubmitSvc struct {
	BasePath           string
	SocketPath         string
	MaxQueueSize       int32
	MaxConcurrentJudge int32
	PJCompilerPrefix   string
}

type DatabaseSvc struct {
	Conn *sql.DB
}

func New(cfg *config.Config) (*Service, error) {
	if cfg == nil {
		return nil, errors.New("no cfg")
	}

	var svc Service

	svc.Debug = cfg.Debug

	/* problem */
	if cfg.Problem.BasePath == "" {
		return nil, errors.New("no problem path")
	}
	svc.Problem.BasePath = cfg.Problem.BasePath
	if cfg.Problem.InCasePrefix != "" {
		svc.Problem.InCasePrefix = cfg.Problem.InCasePrefix
	} else {
		svc.Problem.InCasePrefix = "input"
	}
	if cfg.Problem.InCaseSuffix != "" {
		svc.Problem.InCaseSuffix = cfg.Problem.InCaseSuffix
	} else {
		svc.Problem.InCaseSuffix = ".bin"
	}
	if cfg.Problem.OutCasePrefix != "" {
		svc.Problem.OutCasePrefix = cfg.Problem.OutCasePrefix
	} else {
		svc.Problem.OutCasePreifx = "output"
	}
	if cfg.Problem.OutCaseSuffix != "" {
		svc.Problem.OutCaseSuffix = cfg.Problem.OutCaseSuffix
	} else {
		svc.Problem.OutCaseSuffix = ".bin"
	}

	/* submit */
	if cfg.Submit.BasePath == "" || cfg.Submit.SocketPath == "" {
		return nil, errors.New("not submit/socket path")
	}
	svc.Submit.BasePath = cfg.Submit.BasePath
	svc.Submit.SocketPath = cfg.Submit.SocketPath
	svc.Submit.PJCompilerPrefix = "compiler_type="
	if cfg.Submit.MaxConcurrentJudge > 0 {
		svc.Submit.MaxQueueSize = cfg.Submit.MaxConcurrentJudge
	} else {
		svc.Submit.MaxQueueSize = 128
	}
	if cfg.Submit.MaxConcurrentJudge > 0 {
		svc.Submit.MaxConcurrentJudge = cfg.Submit.MaxConcurrentJudge
	} else {
		svc.Submit.MaxConcurrentJudge = 8
	}

	return &svc
}
