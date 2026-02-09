package service

import (
	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type MenuService struct {
	MenuRepository *dal.MenuRepository
}

func NewMenuService(menuRepository *dal.MenuRepository) *MenuService {
	return &MenuService{MenuRepository: menuRepository}
}

func (s *MenuService) GetAll() ([]models.MenuItem, error) {
	menuItems, err := s.MenuRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return menuItems, nil
}
