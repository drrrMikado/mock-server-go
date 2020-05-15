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

type AddMockParam struct {
	Description string `json:"description" form:"description" binding:"required"`
	Uri         string `json:"uri" form:"uri" binding:"required"`
	Method      string `json:"method" form:"method" binding:"required"`
	Delay       string `json:"delay" form:"delay" binding:"required"`
	StatusCode  int    `json:"status_code" form:"status_code" binding:"required"`
	Headers     string `json:"headers" form:"headers"`
	Body        string `json:"body" form:"body"`
}

type UpdateMockParam struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Uri         string `json:"uri" form:"uri" binding:"required"`
	Method      string `json:"method" form:"method" binding:"required"`
	Delay       string `json:"delay" form:"delay" binding:"required"`
	StatusCode  int    `json:"status_code" form:"status_code" binding:"required"`
	Headers     string `json:"headers" form:"headers"`
	Body        string `json:"body" form:"body"`
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
