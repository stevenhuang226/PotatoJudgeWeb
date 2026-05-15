package service

import (
	"errors"
	"pjweb/internal/config"
)

const (
	SubmissionQueued     string = "queued"
	SubmissionJudging    string = "judging"
	SubmissionDone       string = "done"
	PJCompilerTypePrefix string = "compiler_type="
	DefMaxQueueSize      int    = 128
	DefMaxConcurrentJobs int    = 8
)

type SubmitService struct {
	MaxQueueSize      int
	MaxConcurrentJobs int
}

type Submission struct {
	uint32 sub_id
	uint32 pro_id
	uint32 compilerType
}

func New(cfg *config.SubmitConfig) (*SubmitService, error) {
	if cfg == nil {
		return nil, errors.New("no submit cfg")
	}

	MaxQueueSize := cfg.MaxQueueSize
	if MaxQueueSize <= 0 {
		MaxQueueSize = DefMaxQueueSize
	}

	MaxConcurrentJobs := cfg.MaxConcurrentJudge
	if MaxConcurrentJobs <= 0 {
		MaxConcurrentJobs = DefMaxConcurrentJobs
	}

	var svc SubmitService
	svc.MaxQueueSize = MaxQueueSize
	svc.MaxConcurrentJobs = MaxConcurrentJobs

	return &svc, nil
}

func (svc *SubmitService) Submit(sub *Submission) error {
	return nil
}
