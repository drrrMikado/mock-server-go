package services

import (
	"testing"

	"github.com/drrrMikado/mock-server-go/conf"
	"github.com/stretchr/testify/assert"
)

func init() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	New(conf.Conf)
}

func TestService_GetMockCfg(t *testing.T) {
	_, err := s.GetMockCfg(int64(-1))
	assert.Error(t, err)
	m, err := s.GetMockCfg(int64(1))
	assert.NotNil(t, m)
	assert.Nil(t, err)
}
