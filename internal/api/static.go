package api

import (
	"io"
	"net/http"
	"pjweb/internal/service"
)

type Static struct {
	Service *service.StaticService
}

func NewStatic(staticSvc *service.StaticService) (*Static, error) {
	return &Static{
		Service: staticSvc,
	}, nil
}

func (svc *Static) Handler(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	file, err := svc.Service.CheckAndGetFile(path)

	if err != nil {
		http.NotFound(res, req)
		return
	}

	defer file.Data.Close()
	res.Header().Set("Content-Type", file.ContentType)
	res.Header().Set("Content-Length", file.Size)

	io.Copy(res, file.Data)
}
