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

func (s *Service) AddMock(mp *models.AddMockParam) (m *models.Mock, err error) {
	m = &models.Mock{
		Uri:    mp.Uri,
		Method: mp.Method,
	}
	if err = s.dao.DB.Model(&models.Mock{}).
		Where("`uri` = ?", mp.Uri).
		Where("`method` = ?", mp.Method).
		Assign(map[string]interface{}{
			"description": mp.Description,
			"delay":       mp.Delay,
			"status_code": mp.StatusCode,
			"headers":     mp.Headers,
			"body":        mp.Body,
		}).FirstOrCreate(m).Error; err != nil {
		return
	}
	return
}

func (s *Service) UpdateMock(mp *models.UpdateMockParam) (err error) {
	if err = s.dao.DB.Model(&models.Mock{ID: mp.ID}).
		Update(map[string]interface{}{
			"uri":         mp.Uri,
			"method":      mp.Method,
			"description": mp.Description,
			"delay":       mp.Delay,
			"status_code": mp.StatusCode,
			"headers":     mp.Headers,
			"body":        mp.Body,
		}).Error; err != nil {
		return
	}
	return
}

func (s *Service) DeleteMock(id uint64) (err error) {
	if err = s.dao.DB.Delete(&models.Mock{ID: id}).Error; err != nil {
		return
	}
	return
}
