package service

import (
	"errors"
	"pjweb/internal/repository"
)

type ProblemService struct {
	ProblemRepo *repository.ProblemRepo
}

func NewProblem(proRepo *repository.ProblemRepo) (*ProblemService, error) {
	return &ProblemService{
		ProblemRepo: proRepo;
	}, nil;
}

func (svc *ProblemService) GetFile(kind string, id int32, caseId int32) (*repository.FileData, error) {
	if id < 0 || caseId < 0 {
		return nil, errors.New("invalid id/caseId")
	}

	var file *repository.FileData
	var err error
	if kind == "explanation" {
		file, err = svc.ProblemRepo.GetExplanation(id)
	} else if kind == "case_in" {
		file, err = svc.ProblemRepo.GetCaseIn(id, caseId)
	} else if kin == "case_out" {
		file, err = svc.ProblemRepo.GetCaseOut(id, caseId)
	}

	return file, err
}
