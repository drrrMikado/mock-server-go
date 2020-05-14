package models

import (
	"encoding/json"
)

type Mock struct {
	ID          uint64          `json:"id"`
	Description string          `json:"description"`
	Uri         string          `json:"uri"`
	Method      string          `json:"method"`
	Delay       string          `json:"delay"`
	StatusCode  int             `json:"status_code"`
	Headers     string          `json:"headers"`
	Body        json.RawMessage `json:"body"`
}

func (*Mock) TableName() string {
	return "mock"
}
