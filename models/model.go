package models

import (
	"database/sql/driver"
	"encoding/json"
)

type Mock struct {
	ID          uint64          `json:"id"`
	Description string          `json:"description"`
	Uri         string          `json:"uri"`
	Method      string          `json:"method"`
	Delay       string          `json:"delay"`
	StatusCode  int             `json:"status_code"`
	Headers     Header          `json:"headers"`
	Body        json.RawMessage `json:"body"`
}

func (*Mock) TableName() string {
	return "mock"
}

type Header map[string]string

func (h *Header) Value() (driver.Value, error) {
	b, err := json.Marshal(h)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func (h *Header) Scan(src interface{}) (err error) {
	switch v := src.(type) {
	case []uint8:
		if err = json.Unmarshal(v, &h); err != nil {
			return err
		}
	}
	return
}
