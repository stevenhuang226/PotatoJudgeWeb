package api

import (
	"io"
	"net/http"
	"pjweb/internal/service"
	"strconv"
)

func (svc *service.Service) ProblemHandler(res http.ResponseWriter, req *http.Request) {
	kind := req.URL.Query().Get("kind")
	idStr := req.URL.Query().Get("id")
	caseIdStr := req.URL.Query().Get("case_id")

	id, err := strconv.Atoi(idStr)
	caseId, err := strconv.Atoi(caseIdStr)

	if kind == "explanation" {
		/* is explanation */
		file, err := svc.Problem.GetExplanation(id)
		if err != nil {
			http.NotFound(res, req)
			return
		}
		defer file.Data.Close()

		res.Header().Set("Content-Type", "text/plain")
		res.Header().Set("Content-Length", strconv.FormatInt(file.Size, 10))

		io.Copy(res, file.Data)
	} else if kind == "cases_in" {
		/* case in */
		file, err := svc.Problem.GetInCase(id, caseId)
		if err != nil {
			http.NotFound(res, req)
			return
		}
		defer file.Data.Close()

		res.Header().Set("Content-Type", "application/octet-stream")
		res.Header().Set("Content-Length", strconv.FormatInt(file.Size, 10))

		io.Copy(res, file.Data)
	} else if kind == "case_out" {
		/* case out */
		file, err := svc.Problem.GetOutCase(id, caseId)
		if err != nil {
			http.NotFound(res, req)
			return
		}
		defer file.Data.Close()

		res.Header().Set("Content-Type", "application/octet-stream")
		res.Header().Set("Content-Length", strconv.FormatInt(file.Size, 10))

		io.Copy(res, file.Data)
	} else {
		http.Error(res, "variable issues", http.StatusBadRequest)
	}
}
