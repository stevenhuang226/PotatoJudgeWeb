package app

import (
	"log"
	"net/http"
	"pjweb/internal/api"
	"pjweb/internal/repository"
	"pjweb/internal/service"
	"strconv"
)

func (app *App) Run() error {
	var err error

	app.ProblemRepo, err = repository.NewProblemRepo(
		app.Problem.BasePath,
		app.Problem.Explanation,
		app.Problem.InCasePrefix,
		app.Problem.InCaseSuffix,
		app.Problem.OutCasePrefix,
		app.Problem.OutCaseSuffix,
	)
	if err != nil {
		return err
	}

	app.SubmitRepo, err = repository.NewSubmitRepo(
		app.Submit.BasePath,
		app.Submit.SocketPath,
		app.Submit.MaxQueueSize,
		app.Submit.MaxConcurrentJudge,
	)
	if err != nil && app.SubmitRepo == nil {
		return err
	}

	/* debug */
	if !app.Debug.IsOn || app.Debug.UseDB {
		app.DatabaseRepo, err = repository.NewDatabaseRepo(app.Database.Conn)
		if err != nil {
			return err
		}
	} else {
		app.DatabaseRepo = nil
	}
	/* end debug */

	app.StaticRepo, err = repository.NewStaticRepo(
		app.Static.BasePath,
	)
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

	mux, err := app.Router.NewMux()
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(int(app.Server.Port)),
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())

	if !app.Debug.IsOn || app.Debug.UseDB {
		app.Database.Conn.Close()
	}

	return nil
}
