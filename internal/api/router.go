package api

import (
	"net/http"
	"pjweb/internal/config"
)

type Service struct {
	Problem ProblemService
	Static  StaticService
}

func NewRouter(svc *Service) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/problem", svc.Problem.Handler)
	mux.HandleFunc("/", svc.Static.Handler)

	return mux
}

func RouterServiceByConfig(cfg *config.Config) (*Service, error) {
	var svc Service

	svc.Problem.BasePath = cfg.Problem.BasePath
	svc.Problem.CasePrefix = cfg.Problem.CasePrefix
	svc.Problem.CaseSuffix = cfg.Problem.CaseSuffix
	svc.Problem.ExplanationName = cfg.Problem.ExplanationName

	svc.Static.BasePath = cfg.Static.BasePath

	return &svc, nil
}
