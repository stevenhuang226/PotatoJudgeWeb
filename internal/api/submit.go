package api

import (
	"errors"
	"io"
	"net/http"
	"pjweb/internal/service"
	"strconv"
)

const (
	MaxRequestSize  int64 = 64 * 1024 // 64Kb
	MinCompilerType       = 1
	MaxCompilerType       = 2
)

type Submit struct {
	Service *service.SubmitService
}

func NewSubmit(subSvc *service.SubmitService) (*Submit, error) {
	return &Submit{
		Service: subSvc,
	}, nil
}

/*
the url should be like
domain/submit?problem_id=<id>&compiler_type=<id>
*/

func (svc *Submit) Handler(res http.ResponseWriter, req *http.Request) {
	req.Body = http.MaxBytesREader(res, req.Body, svc.MaxRequestSize)

	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			http.Error(
				res,
				"request too large",
				http.StatusRequestEntityTooLarge,
			)
		} else {
			http.Error(
				res,
				"error",
				http.StatusBadRequest,
			)
		}
		return
	}

	problemIdStr := req.URL.Query().Get("problem_id")
	problemId, err := strconv.Atoi(problemIdStr)
	if err != nil {
		http.Error(res, "invalid id", http.StatusBadRequest)
		return
	}

	compilerTypeStr := req.URL.Query().Get("compiler_type")
	compilerType, err := strconv.Atoi(compilerTypeStr)
	if err != nil {
		http.Error(res, "invalid compiler type", http.StatusBadRequest)
		return
	}

	err := svc.Service.Submit(problemId, compilerType, string(reqBody))

	if err != nil {
		http.Error(res, "submit failed", http.StatusBadRequest)
		return
	}

	res.Write([]byte("submit success"))
}
