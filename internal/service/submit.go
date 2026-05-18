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
	BasePath          string
	SocketPath        string
	MaxQueueSize      int
	MaxConcurrentJobs int
}

type Submission struct {
	Id           uint32
	ProblemId    uint32
	CompilerType int32
	Source       string
}

func New(cfg *config.SubmitConfig) (*SubmitService, error) {
	if cfg == nil {
		return nil, errors.New("no submit cfg")
	}

	basePath := cfg.BasePath
	if basePath == "" {
		return nil, errors.New("no submit base path")
	}

	socketPath := cfg.SocketPath
	if socketPath == "" {
		return nil, errors.New("no submit socket")
	}

	maxQueueSize := cfg.MaxQueueSize
	if maxQueueSize <= 0 {
		maxQueueSize = DefMaxQueueSize
	}

	maxConcurrentJobs := cfg.MaxConcurrentJudge
	if maxConcurrentJobs <= 0 {
		maxConcurrentJobs = DefMaxConcurrentJobs
	}

	var svc SubmitService
	svc.MaxQueueSize = maxQueueSize
	svc.MaxConcurrentJobs = maxConcurrentJobs
	svc.BasePath = basePath
	svc.SocketPath = socketPath

	return &svc, nil
}

func (svc *SubmitService) Submit(sub *Submission) error {
	return nil

	/* check problem exist */
}
