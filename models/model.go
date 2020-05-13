package models

import (
	"github.com/drrrMikado/mock-server-go/internal/time"
)

type Mock struct {
	ID         string            `json:"id,omitempty"`
	Desc       string            `json:"description"`
	Uri        string            `json:"uri"`
	Method     string            `json:"method"`
	Delay      time.Duration     `json:"delay"`
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}
