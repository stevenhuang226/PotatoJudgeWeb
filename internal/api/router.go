package api

import (
	"errors"
	"net/http"
	"pjweb/internal/service"
)

type Router struct {
	Problem Problem
	Submit  Submit
	Static  Static
}

func NewMux(router *Router) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/prob", router.Problem.Handler)
	mux.HandleFunc("/submit", router.Submit.Handler)
	m

	mux.HandleFunc("/problem", svc.Problem.Handler)
	mux.HandleFunc("/", svc.Static.Handler)

	return mux
}

func NewRouter(
	proSvc *service.ProblemService,
	subSvc *service.SubmitService,
	staticSvc *service.StaticService,
) (*Router, error) {
	if proSvc == nil || subSvc == nil || staticSvc == nil {
		return nil, errors.New("miss service")
	}

	return &Router{
		Problem: NewProblem(proSvc),
		Submit:  NewSubmit(subSvc),
		Static:  NewStatic(staticSvc),
	}, nil
}
