package api

import "net/http"

type SubmitService struct {
	BasePath           string
	SocketPath         string
	MaxConcurrentJudge uint16
}

func (svc *SubmitService) Handler(res http.ResponseWriter, req *http.Request) {
	/*
		read use submit and send to judge system
		requets to check current run state. Avoid sending too much same time. <= requets DB use
	*/
}
