package api

import (
	"errors"
	"net/http"
	"pjweb/internal/service"
)

type Router struct {
	Problem *Problem
	Submit  *Submit
	Static  *Static
}

func (router *Router) NewMux() (http.Handler, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/prob", router.Problem.Handler)
	mux.HandleFunc("/submit", router.Submit.Handler)

	mux.HandleFunc("/problem", router.Problem.Handler)
	mux.HandleFunc("/", router.Static.Handler)

	return mux, nil
}

func NewRouter(
	proSvc *service.ProblemService,
	subSvc *service.SubmitService,
	staticSvc *service.StaticService,
) (*Router, error) {
	if proSvc == nil || subSvc == nil || staticSvc == nil {
		return nil, errors.New("miss service")
	}

	var router Router
	var err error

	router.Problem, err = NewProblem(proSvc)
	if err != nil {
		return nil, err
	}

	router.Submit, err = NewSubmit(subSvc)
	if err != nil {
		return nil, err
	}

	router.Static, err = NewStatic(staticSvc)
	if err != nil {
		return nil, err
	}

	return &router, nil
}
