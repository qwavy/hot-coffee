package service

import "hot-coffee/internal/dal"

type MenuService struct {
	MenuRepository dal.MenuRepository
}

func NewMenuService(menuRepository dal.MenuRepository) *MenuService {
	return &MenuService{MenuRepository: menuRepository}
}
