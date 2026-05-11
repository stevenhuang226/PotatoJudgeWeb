package api

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type ProblemService struct {
	BasePath        string
	ExplanationName string
	CasePrefix      string
	CaseSuffix      string
}

/*
the url will be like:
domain/problem?id=<id>&kind=explanation
domain/problem?id=<id>&kind=test_cases&case_id=<case_id>
*/

func (svc *ProblemService) Handler(res http.ResponseWriter, req *http.Request) {
	const (
		KindExplanation string = "explanation"
		KindTestCases   string = "test_cases"
		KindErrorMsg    string = "invalid kind"
	)

	kind := req.URL.Query().Get("kind")

	switch kind {
	case KindExplanation:
		svc.ExplanationHandler(res, req)
	case KindTestCases:
		svc.TestCasesHandler(res, req)
	default:
		http.NotFound(res, req)
		return
	}
}

func (svc *ProblemService) ExplanationHandler(res http.ResponseWriter, req *http.Request) {
	const (
		ExplanationFileName string = "explanation.md"
	)

	idStr := req.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.NotFound(res, req)
		return
	}

	target := filepath.Join(svc.BasePath, strconv.Itoa(id), svc.ExplanationName)

	data, err := os.ReadFile(target)
	if err != nil {
		http.NotFound(res, req)
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.Write(data)

	return
}

func (svc *ProblemService) TestCasesHandler(res http.ResponseWriter, req *http.Request) {
	idStr := req.URL.Query().Get("id")
	caseIdStr := req.URL.Query().Get("case_id")

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.NotFound(res, req)
		return
	}

	caseId, err := strconv.Atoi(caseIdStr)
	if err != nil || caseId < 0 {
		http.NotFound(res, req)
		return
	}

	fileName := svc.CasePrefix + strconv.Itoa(caseId) + svc.CaseSuffix
	target := filepath.Join(svc.BasePath, strconv.Itoa(id), fileName)

	data, err := os.ReadFile(target)
	if err != nil {
		http.NotFound(res, req)
		return
	}

	res.Header().Set("Content-Type", "application/octet-stream")
	res.Write(data)

	return
}
