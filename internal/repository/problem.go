package repository

import (
	"os"
	"path/filepath"
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

func NewProblemRepo(
	basePath string,
	explanation string,
	caseInPrefix string,
	caseInSuffix string,
	caseOutPrefix string,
	caseOutSuffix string,
) (*ProblemRepo, error) {
	return &ProblemRepo{
		BasePath:      basePath,
		Explanation:   explanation,
		CaseInPrefix:  caseInPrefix,
		CaseInSuffix:  caseInSuffix,
		CaseOutPrefix: caseOutPrefix,
		CaseOutSuffix: caseOutSuffix,
	}, nil
}

func (repo *ProblemRepo) GetExplanation(id int) (*FileData, error) {
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

func (repo *ProblemRepo) GetCaseIn(id int, caseId int) (*FileData, error) {
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
		Data:        file,
		Size:        stat.Size(),
	}, nil
}

func (repo *ProblemRepo) GetCaseOut(id int, caseId int) (*FileData, error) {
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
		Data:        file,
		Size:        stat.Size(),
	}, nil
}

func (repo *ProblemRepo) DirExist(id uint) bool {
	target := filepath.Join(
		repo.BasePath,
		strconv.Itoa(int(id)),
	)

	stat, err := os.Stat(target)

	if err != nil {
		return false
	}

	return stat.IsDir()
}
