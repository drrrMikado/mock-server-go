package services

import (
	"errors"

	"github.com/drrrMikado/mock-server-go/models"
)

func (s *Service) GetMock(id int) (m *models.Mock, err error) {
	if id <= 0 {
		err = errors.New("id must greater than zero")
		return
	}
	m = &models.Mock{}
	err = s.dao.DB.First(m, id).Error
	return
}

func (s *Service) GetMocks(page, pageSize int) (mlp *models.MockListPage, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize must greater than zero")
		return
	} else if pageSize > 100 {
		err = errors.New("pageSize must lesser than 100")
		return
	}
	var (
		ms    []*models.Mock
		count int
	)
	s.dao.DB.Model(&models.Mock{}).Count(&count)
	if err = s.dao.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&ms).Error; err != nil {
		return
	}
	mlp = &models.MockListPage{
		Items: ms,
		Page: &models.Page{
			PageSize: pageSize,
			Page:     page,
			Total:    count,
		},
	}
	return
}

func (s *Service) GetMockByUriAndMethod(uri, method string) (m *models.Mock, err error) {
	if len(uri) == 0 || len(method) == 0 {
		err = errors.New("uri or method must has value")
		return
	}
	m = &models.Mock{}
	err = s.dao.DB.Where("`uri` = ?", uri).
		Where("`method` = ?", method).
		First(m).Error
	return
}
