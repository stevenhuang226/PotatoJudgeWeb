package repository

import "io"

type FileData struct {
	Name        string
	ContentType string
	Data        io.ReadCloser
	Size        int64
}
