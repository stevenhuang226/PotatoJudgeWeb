package repository

import (
	"errors"
	"mime"
	"os"
	"path/filepath"
	"pjweb/internal/app"
)

type StaticRepo struct {
	BasePath string
}

func NewStaticRepo(static *app.Static) (*StaticRepo, error) {
	if static == nil {
		return nil, errors.New("no static")
	}

	return &StaticRepo{
		BasePath: static.BasePath,
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
