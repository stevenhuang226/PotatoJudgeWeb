package api

import (
	"io"
	"net/http"
	"pjweb/internal/repository"
	"pjweb/internal/service"
	"strconv"
)

type Problem struct {
	Service *service.ProblemService
}

func NewProblem(proSvc *service.ProblemService) (*Problem, error) {
	return &Problem{
		Service: proSvc,
	}, nil
}

func (svc *Problem) Handler(res http.ResponseWriter, req *http.Request) {
	kind := req.URL.Query().Get("kind")
	idStr := req.URL.Query().Get("id")
	caseIdStr := req.URL.Query().Get("case_id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(
			res,
			"invalid id",
			http.StatusBadRequest,
		)
		return
	}

	caseId, err := strconv.Atoi(caseIdStr)
	if caseIdStr != "" && err != nil {
		http.Error(
			res,
			"invalid case id",
			http.StatusBadRequest,
		)
		return
	}

	var file *repository.FileData

	file, err = svc.Service.GetFile(kind, id, caseId)

	if err != nil {
		http.NotFound(res, req)
		return
	}

	defer file.Data.Close()

	res.Header().Set("Content-Type", file.ContentType)
	res.Header().Set("Content-Length", strconv.Itoa(int(file.Size)))

	io.Copy(res, file.Data)
}
