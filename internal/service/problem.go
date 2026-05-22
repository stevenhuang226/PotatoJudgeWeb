package service

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

type ProblemFile struct {
	Name        string
	ContentType string
	Data        io.ReadCloser
	Size        int64
}

func (pro *ProblemSvc) GetExplanation(id int32) (*ProblemFile, error) {
	if id < 0 {
		return nil, errors.New("invalid id")
	}
	target := filepath.Join(
		pro.BasePath,
		strconv.Itoa(id),
		pro.Explanation,
	)

	f, err := os.Open(target)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, err
	}

	return &ProblemFile{
		Name:        stat.Name(),
		ContentType: "text/plain",
		Data:        f,
		Size:        stat.Size(),
	}, nil
}

func (pro *ProblemSvc) GetInCase(id int32, caseId int32) (*ProblemFile, error) {
	if id < 0 || caseId < 0 {
		return nil, errors.New("invalid id")
	}

	target := filepath.Join(
		pro.BasePath,
		strconv.Itoa(id),
		pro.InCasePrefix+strconv.Itoa(caseId)+pro.InCaseSuffix,
	)

	f, err := os.Open(target)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, err
	}

	return &ProblemFile{
		Name:        stat.Name(),
		ContentType: "application/octet-stream",
		Data:        f,
		Size:        stat.Size(),
	}, nil
}

func (pro *ProblemSvc) GetOutCase(id int32, caseId int32) (*ProblemFile, error) {
	if id < 0 || caseId < 0 {
		return nil, errors.New("invalid id")
	}

	target := filepath.Join(
		pro.BasePath,
		strconv.Iota(id),
		pro.OutCasePrefix+strconv.Iota(caseId)+pro.OutCaseSuffix,
	)

	f, err := os.Open(target)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, err
	}

	return &ProblemFile{
		Name:        stat.Name(),
		ContentType: "application/octet-stream",
		Data:        f,
		Size:        stat.Size(),
	}, nil
}
