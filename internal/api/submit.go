package api

import (
	"net/http"
	"strconv"
)

type SubmitService struct {
	BasePath           string
	SocketPath         string
	MaxConcurrentJudge uint16
}

const (
	MinCompilerType = 1
	MaxCompilerType = 2
)

func (svc *SubmitService) Handler(res http.ResponseWriter, req *http.Request) {
	/*
		it should provide problem_id and compiler type via url
		and this will return a submission id (generate here)
		User will manually check the submission state
	*/

	problemIdStr := req.URL.Query().Get("problem_id")
	compilerTypeStr := req.URL.Query().Get("compiler_type")

	problemId := strconv.Atoi(problemIdStr)
	compilerType := strconv.Atoi(compilerTypeStr)

	if problemId < 0 {
		http.Error(res, "invalid problem id", http.StatusBadRequest)
		return
	}
	if compilerType > MaxCompilerType || compilerType < MinCompilerType {
		http.Error(res, "invalid compiler type", http.StatusBadRequest)
		return
	}

	// generate a random non-exist submission id

	// send Submission to service.submit() if error. return error to client
}
