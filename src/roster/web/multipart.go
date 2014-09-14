package web

import (
	"fmt"
	"github.com/gocraft/web"
	"io"
)

const (
	_        = iota // ignore first value by assigning to blank identifier
	KB int64 = 1 << (10 * iota)
	MB
	GB
)

type MultipartContext struct {
	// mailer *smtp.Client
	maxFormMemory int64 // Default 10 megabytes
}

func (c *MultipartContext) MaxFormMemory() int64 {
	if c.maxFormMemory == 0 {
		return 10 * MB
	}
	return c.maxFormMemory
}

func (c *MultipartContext) GetSingleFile(r *web.Request, filename string) (io.ReadCloser, error) {
	err := r.ParseMultipartForm(c.MaxFormMemory())
	if err != nil {
		return nil, err
	}
	m := r.MultipartForm
	header := m.File[filename]
	if len(header) != 1 {
		return nil, fmt.Errorf("You must upload one file")
	}
	f, err := header[0].Open()
	if err != nil {
		return nil, err
	}
	return f, nil
}
