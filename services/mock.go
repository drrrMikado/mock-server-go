package services

import (
	"errors"

	"github.com/drrrMikado/mock-server-go/models"
)

func (s *Service) GetMockCfg(id int64) (m *models.Mock, err error) {
	if id <= 0 {
		err = errors.New("id must greater than zero")
		return
	}
	m = &models.Mock{}
	err = s.dao.DB.First(m, id).Error
	return
}
