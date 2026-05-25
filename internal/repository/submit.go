package repository

import (
	"encoding/binary"
	"errors"
	"net"
	"pjweb/internal/app"
)

type SubmitRepo struct {
	BasePath           string
	SocketPath         string
	MaxQueueSize       int32
	MaxConcurrentJudge int32
	PJCompilerPrefix   string
	PJDetailName       string
	PJConn             net.Conn
}

type PJCSubmissionStruct struct {
	submission_id uint32
	problem_id    uint32
}

func NewSubmit(sub *app.Submit) (*SubmitRepo, error) {
	var repo SubmitRepo

	repo.BasePath = sub.BasePath
	repo.SocketPath = sub.SocketPath
	repo.MaxQueueSize = sub.MaxQueueSize
	repo.MaxConcurrentJudge = sub.MaxConcurrentJudge
	repo.PJCompilerPrefix = "compiler_type="
	repo.PJDetailName = "detail.conf"

	err := repo.InitPJConn()
	if err != nil {
		return &repo, errors.New("init pj socket conn failed")
	}

	return &repo, nil
}

func (repo *SubmitRepo) InitPJConn() error {
	socketPath := repo.SocketPath

	conn, err := net.Dial("unix", socketPath)

	if err != nil {
		repo.PJConn = nil
		return err
	}

	repo.PJConn = conn
	return nil
}

func (repo *SubmitRepo) SendToPJ(subId int, proId int) error {
	if subId < 0 || proId < 0 {
		return errors.New("invalid sub/pro Id")
	}
	if repo.PJConn == nil {
		return errors.New("no pj conn")
	}

	msg := PJCSubmissionStruct{
		problem_id:    uint32(proId),
		submission_id: uint32(subId),
	}

	err := binary.Write(repo.PJConn, binary.LittleEndian, msg)

	return err
}
