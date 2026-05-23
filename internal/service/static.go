package service

import (
	"errors"
	"path/filepath"
	"pjweb/internal/repository"
	"strings"
)

type StaticService struct {
	StaticRepo *repository.StaticRepo
}

func NewStatic(static *repository.StaticRepo) (*StaticService, error) {
	if static == nil {
		return nil, errors.New("no static repo")
	}

	return &StaticService{
		StaticRepo: static,
	}, nil
}

func (svc *StaticService) CheckAndGetFile(userPath string) (*repository.FileData, error) {
	baseAbsPath := filepath.Abs(svc.StaticRepo.BasePath)

	cleanPath := filepath.Clean(userPath)
	fullAbsPath := filepath.Join(baseAbsPath, cleanPath)

	if !strings.HasPrefix(fullAbsPath, baseAbsPath) {
		return nil, errors.New("wrong path")
	}

	return svc.StaticRepo.GetFile(fullAbsPath)
}
