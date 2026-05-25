package app

import (
	"errors"
	"net/http"
	"pjweb/internal/api"
	"pjweb/internal/repository"
	"pjweb/internal/service"
	"strconv"
)

func (app *App) Run() error {
	app.ProblemRepo, err = repository.NewProblem(app.Problem)
	if err != nil {
		return err
	}
	app.SubmitRepo, err = repository.NewSubmit(app.Submit)
	if err != nil {
		return err
	}
	app.DatabaseRepo, err = repository.NewDatabase(app.Database)
	if err != nil {
		return err
	}
	app.StaticRepo, err = repository.NewStatic(app.Static)
	if err != nil {
		return err
	}

	app.ProblemSvc, err = service.NewProblem(app.ProblemRepo)
	if err != nil {
		return err
	}
	app.SubmitSvc, err = service.NewSubmit(app.SubmitRepo, app.ProblemRepo, app.DatabaseRepo)
	if err != nil {
		return err
	}
	app.StaticSvc, err = service.NewStatic(app.StaticRepo)
	if err != nil {
		return err
	}

	app.Router, err = api.NewRouter(
		app.ProblemSvc,
		app.SubmitSvc,
		app.StaticSvc,
	)
	if err != nil {
		return err
	}

	mux := app.Router.NewMux()
	if mux == nil {
		return errors.New("no mux")
	}

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(int(app.Server.Port)),
		Handler: mux,
	}

	server.ListenAndServe()
	app.Database.Conn.Close()
}
