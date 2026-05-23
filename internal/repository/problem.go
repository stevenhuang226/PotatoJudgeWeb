package repository

import (
	"os"
	"path/filepath"
	"pjweb/internal/app"
	"strconv"
)

type ProblemRepo struct {
	BasePath      string
	Explanation   string
	CaseInPrefix  string
	CaseInSuffix  string
	CaseOutPrefix string
	CaseOutSuffix string
}

func NewProblem(pro *app.Problem) (*ProblemRepo, error) {
	var repo ProblemRepo

	repo.BasePath = pro.BasePath
	repo.Explanation = pro.Explanation
	repo.CaseInPrefix = pro.InCasePrefix
	repo.CaseInSuffix = pro.InCaseSuffix
	repo.CaseOutPrefix = pro.OutCasePrefix
	repo.CaseOutSuffix = pro.OutCaseSuffix

	return &repo, nil
}

func (repo *ProblemRepo) GetExplanation(id int32) (*FileData, error) {
	target := filepath.Join(
		repo.BasePath,
		strconv.Itoa(id),
		repo.Explanation,
	)

	file, err := os.Open(target)
	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, err
	}

	return &FileData{
		Name:        stat.Name(),
		ContentType: "text/plain",
		Data:        file,
		Size:        stat.Size(),
	}, nil
}

func (repo *ProblemRepo) GetCaseIn(id int32, caseId int32) (*FileData, error) {
	target := filepath.Join(
		repo.BasePath,
		strconv.Itoa(id),
		repo.CaseInPrefix+strconv.Itoa(caseId)+repo.CaseInSuffix,
	)

	file, err := os.Open(target)
	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, err
	}

	return &FileData{
		Name:        stat.Name(),
		ContentType: "application/octet-stream",
		File:        file,
		Size:        stat.Size(),
	}, nil
}

func (repo *ProblemRepo) GetCaseOut(id int32, caseId int32) (*FileData, error) {
	target := filepath.Join(
		repo.BasePath,
		strconv.Itoa(id),
		repo.CaseOutPrefix+strconv.Itoa(caseId)+repo.CaseOutSuffix,
	)

	file, err := os.Open(target)
	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, err
	}

	return &FileData{
		Name:        stat.Name(),
		ContentType: "application/octet-stream",
		File:        file,
		Size:        stat.Size(),
	}, nil
}

func (repo *ProblemRepo) DirExist(id int32) bool {
	target := filepath.Join(
		repo.BasePath,
		strconv.Itoa(id),
	)

	stat, err := os.Stat(target)

	if err != nil {
		return false
	}

	return stat.IsDir()
}
