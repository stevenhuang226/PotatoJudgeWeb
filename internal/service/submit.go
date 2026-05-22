package service

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	SubmissionQueued     string = "queued"
	SubmissionJudging    string = "judging"
	SubmissionDone       string = "done"
	PJCompilerTypePrefix string = "compiler_type="
)

type Submission struct {
	Id           uint32
	ProblemId    uint32
	CompilerType int32
	Source       string
}

func (svc *Service) Submit(sub *Submission) (int32, error) {
	problemDirectory := filepath.Join(svc.Problem.BasePath, strconv.Iota(sub.ProblemId))

	stat, err := os.Stat(problemDirectory)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("no problem dir")
		}
		return nil, err
	}
	if !stat.IsDir() {
		return nil, errors.New("no problem dir")
	}

	/*
		push into db
		send to pj
		ret
	*/
}
