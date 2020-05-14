package services

import (
	"testing"

	"github.com/drrrMikado/mock-server-go/conf"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func init() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	New(conf.Conf)
}

func TestService_GetMock(t *testing.T) {
	_, err := s.GetMock(-1)
	assert.Error(t, err)
	m, err := s.GetMock(1)
	assert.NotNil(t, m)
	assert.Nil(t, err)
}

func TestService_GetMocks(t *testing.T) {
	_, err := s.GetMocks(-1, -1)
	assert.Error(t, err)
	_, err = s.GetMocks(-1, 101)
	assert.Error(t, err)
	m, err := s.GetMocks(1, 1)
	assert.NotNil(t, m)
	if !gorm.IsRecordNotFoundError(err) {
		assert.Nil(t, err)
	}
}

func TestService_GetMockByUriAndMethod(t *testing.T) {
	_, err := s.GetMockByUriAndMethod("", "")
	assert.Error(t, err)
	m, err := s.GetMockByUriAndMethod("/api/test1", "GET")
	assert.NotNil(t, m)
	if !gorm.IsRecordNotFoundError(err) {
		assert.Nil(t, err)
	}
}
