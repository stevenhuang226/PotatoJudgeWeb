package repository

import (
	"encoding/binary"
	"errors"
	"net"
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

func NewSubmitRepo(
	basePath string,
	socketPath string,
	maxQueueSize int32,
	maxConcurrentJudge int32,
) (*SubmitRepo, error) {
	repo := SubmitRepo{
		BasePath:           basePath,
		SocketPath:         socketPath,
		MaxQueueSize:       maxQueueSize,
		MaxConcurrentJudge: maxConcurrentJudge,
		PJCompilerPrefix:   "compiler_type=",
		PJDetailName:       "detail.conf",
	}

	err := repo.InitPJConn()
	if err != nil {
		return &repo, errors.New("init pj conn failed")
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
