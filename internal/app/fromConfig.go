package app

import (
	"errors"
	"pjweb/internal/config"
)

func (app *App) SetFromConfig(cfg *config.Config) error {
	if cfg == nil {
		return errors.New("no config")
	}

	app.Debug = cfg.Debug

	/* problem */

	app.Problem.Explanation = "explanation"
	app.Problem.InCasePrefix = "input"
	app.Problem.InCaseSuffix = ".bin"
	app.Problem.OutCasePrefix = "output"
	app.Problem.OutCaseSuffix = ".bin"

	if cfg.Problem.BasePath == "" {
		return errors.New("no problem base")
	}

	app.Problem.BasePath = cfg.Problem.BasePath
	if cfg.Problem.InCasePrefix != "" {
		app.Problem.InCasePrefix = cfg.Problem.InCasePrefix
	}
	if cfg.Problem.InCaseSuffix != "" {
		app.Problem.InCaseSuffix = cfg.Problem.InCaseSuffix
	}
	if cfg.Problem.OutCasePrefix != "" {
		cfg.Problem.OutCasePrefix = cfg.Problem.OutCasePrefix
	}
	if cfg.Problem.OutCaseSuffix != "" {
		cfg.Problem.OutCaseSuffix = cfg.Problem.OutCaseSuffix
	}

	/* submit */

	app.Submit.MaxQueueSize = 128
	app.Submit.MaxConcurrentJudge = 8
	app.Submit.PJCompilerTypePrefix = "compiler_type"
	app.Submit.PJDetailName = "detail.conf"

	if cfg.Submit.BasePath == "" {
		return errors.New("no submit base")
	}

	app.Submit.BasePath = cfg.Submit.BasePath
	if cfg.Submit.MaxQueueSize != 0 {
		app.Submit.MaxQueueSize = cfg.Submit.MaxQueueSize
	}
	if cfg.Submit.MaxConcurrentJudge != 0 {
		app.Submit.MaxConcurrentJudge = cfg.Submit.MaxConcurrentJudge
	}
	if cfg.Submit.PJCompilerTypePrefix != "" {
		app.Submit.PJCompilerTypePrefix = cfg.Submit.PJCompilerTypePrefix
	}
	if cfg.Submit.PJDetailName != "" {
		app.Submit.PJDetailName = cfg.Submit.PJDetailName
	}

	/* static */

	if cfg.Static.BasePath == "" {
		return errors.New("no static base")
	}

	app.Static.BasePath = cfg.Static.BasePath

	return nil
}
