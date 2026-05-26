package repository

import (
	"mime"
	"os"
	"path/filepath"
)

type StaticRepo struct {
	BasePath string
}

func NewStaticRepo(
	basePath string,
) (*StaticRepo, error) {
	return &StaticRepo{
		BasePath: basePath,
	}, nil
}

func (repo *StaticRepo) GetFile(absPath string) (*FileData, error) {
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(absPath)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return &FileData{
		Name:        stat.Name(),
		ContentType: contentType,
		Data:        file,
		Size:        stat.Size(),
	}, nil
}
