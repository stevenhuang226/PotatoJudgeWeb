package service

import (
	"errors"
	"pjweb/internal/repository"
)

const (
	SubmissionQueued     string = "queued"
	SubmissionJudging    string = "judging"
	SubmissionDone       string = "done"
	PJCompilerTypePrefix string = "compiler_type="
)

type SubmitService struct {
	ProblemRepo    *repository.ProblemRepo
	SubmitRepo     *repository.SubmitRepo
	DatabaseRepo   *repository.DatabaseRepo
	MaxRequestSize int64
}

type Submission struct {
	Id           int
	ProblemId    int
	CompilerType int32
	Source       string
}

func NewSubmit(subRepo *repository.SubmitRepo, proRepo *repository.ProblemRepo, dbRepo *repository.DatabaseRepo) (*SubmitService, error) {
	if subRepo == nil || proRepo == nil || dbRepo == nil {
		return nil, errors.New("miss repo")
	}

	return &SubmitService{
		ProblemRepo:    proRepo,
		SubmitRepo:     subRepo,
		DatabaseRepo:   dbRepo,
		MaxRequestSize: 64 * 1024, // 64kb
	}, nil
}

func (svc *SubmitService) Submit(problem_id int, compiler_type int, source string) error {
	if !svc.ProblemRepo.DirExist(problem_id) {
		return errors.New("no problem dir")
	}
}
