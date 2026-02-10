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

func (s *MenuService) GetById(productId string) (models.MenuItem, error) {
	menuItem, err := s.MenuRepository.GetById(productId)

	if err != nil {
		return models.MenuItem{}, err
	}

	return menuItem, err
}

func (s *MenuService) DeleteById(productId string) error {
	return s.MenuRepository.DeleteById(productId)
}

func (s *MenuService) CreateItem(menuItem models.MenuItem) error {
	return s.MenuRepository.CreateItem(menuItem)
}
