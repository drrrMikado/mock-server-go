package services

import (
	"github.com/drrrMikado/mock-server-go/conf"
	"github.com/drrrMikado/mock-server-go/dao"
)

var (
	s *Service
)

type Service struct {
	c   *conf.Config
	dao *dao.Dao
	// waiter      sync.WaitGroup
}

func New(c *conf.Config) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}
}
