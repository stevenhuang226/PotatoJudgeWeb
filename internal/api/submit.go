package api

import (
	"errors"
	"io"
	"net/http"
	"pjweb/internal/service"
	"strconv"
)

type SubmitHandler struct {
	SubmitService *service.SubmitService
}

const (
	MaxRequestSize  int64 = 64 * 1024 // 64Kb
	MinCompilerType       = 1
	MaxCompilerType       = 2
)

/*
the url should be like
domain/submit?problem_id=<id>&compiler_type=<id>
*/
func (handler *SubmitHandler) Handler(res http.ResponseWriter, req *http.Request) {
	req.Body = http.MaxBytesReader(res, req.Body, MaxRequestSize)

	requestBody, err := io.ReadAll(req.Body)
	if err != nil {
		var maxBytesErr *http.MaxBytesError

		if errors.As(err, &maxBytesErr) {
			http.Error(res, "request too large", http.StatusRequestEntityTooLarge)
		} else {
			http.Error(res, "ERROR", http.StatusBadRequest)
		}
		return
	}

	problemIdStr := req.URL.Query().Get("problem_id")
	compilerTypeStr := req.URL.Query().Get("compiler_type")

	problemId, err := strconv.Atoi(problemIdStr)
	if err != nil || problemId < 0 {
		http.Error(res, "invalid problem id", http.StatusBadRequest)
		return
	}

	compilerType, err := strconv.Atoi(compilerTypeStr)
	if err != nil || compilerType > MaxCompilerType || compilerType < MinCompilerType {
		http.Error(res, "invalid compiler type", http.StatusBadRequest)
		return
	}

	var submission service.Submission
	submission.CompilerType = compilerType
	submission.ProblemId = problemId
	submission.Source = string(requestBody)

	err := handler.SubmitService.Submit(&submission)
	if err != nil {
		http.Error(res, "submit failed", http.StatusInternalServerError)
		return
	}

	res.Write([]byte("ok")) // auto set header 200.Ok
	return

	// write submission success back. And Submit(&submission) then exit
}
