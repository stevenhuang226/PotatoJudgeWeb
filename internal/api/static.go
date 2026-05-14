package api

import (
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type StaticService struct {
	BasePath string
}

func (svc *StaticService) Handler(res http.ResponseWriter, req *http.Request) {
	userPath := strings.TrimPrefix(req.URL.Path, svc.BasePath)
	cleanPath := filepath.Clean(userPath)
	fullPath := filepath.Join(svc.BasePath, cleanPath)

	baseAbs, _ := filepath.Abs(svc.BasePath)
	fullAbs, _ := filepath.Abs(fullPath)

	if !strings.HasPrefix(fullAbs, baseAbs) {
		http.NotFound(res, req)
		return
	}

	fileHandle, err := os.Open(fullAbs)
	if err != nil {
		http.NotFound(res, req)
		return
	}
	defer fileHandle.Close()

	ext := filepath.Ext(fullAbs)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	res.Header().Set("Content-Type", contentType)
	res.WriteHeader(http.StatusOK)

	io.Copy(res, fileHandle)
}
