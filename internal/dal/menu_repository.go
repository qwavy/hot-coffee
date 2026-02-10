package dal

import (
	"encoding/json"
	"hot-coffee/models"
	"os"
)

type MenuRepository struct {
	filePath string
}

func NewMenuRepository(filePath string) *MenuRepository {
	return &MenuRepository{filePath: filePath}
}

func (r *MenuRepository) list() ([]models.MenuItem, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var menuItems []models.MenuItem

	err = json.Unmarshal(data, &menuItems)
	if err != nil {
		return nil, err
	}

	return menuItems, nil
}

func (r *MenuRepository) GetAll() ([]models.MenuItem, error) {
	return r.list()
}

func (r *MenuRepository) GetById(productId string) (models.MenuItem, error) {
	menuItems, err := r.list()

	if err != nil {
		return models.MenuItem{}, err
	}

	for _, menuItem := range menuItems {
		if menuItem.ID == productId {
			return models.MenuItem{}, nil
		}
	}

	return models.MenuItem{}, models.MenuItemNotFound
}
