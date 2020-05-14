package models

import (
	"encoding/json"
)

type Mock struct {
	ID          uint            `json:"id"`
	Description string          `json:"description"`
	Uri         string          `json:"uri"`
	Method      string          `json:"method"`
	Delay       string          `json:"delay"`
	StatusCode  int             `json:"status_code"`
	Headers     string          `json:"headers"`
	Body        json.RawMessage `json:"body"`
}

type AddMockParam struct {
	Description string `json:"description"`
	Uri         string `json:"uri"`
	Method      string `json:"method"`
	Delay       string `json:"delay"`
	StatusCode  int    `json:"status_code"`
	Headers     string `json:"headers"`
	Body        string `json:"body"`
}

type Page struct {
	PageSize int `json:"page_size"`
	Page     int `json:"page"`
	Total    int `json:"total"`
}

// BucketListPage bucket/list result
type MockListPage struct {
	Items []*Mock `json:"items"`
	Page  *Page   `json:"page"`
}

func (*Mock) TableName() string {
	return "mock"
}
