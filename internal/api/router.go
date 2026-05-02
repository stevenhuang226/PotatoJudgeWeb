package api

import (
	"net/http"
)

type Service struct {
	Problem ProblemService
}

func NewRouter(svc *Service) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/problem", svc.Problem.Handler)
}
