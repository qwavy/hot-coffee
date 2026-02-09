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

func (r *MenuRepository) GetAll() ([]models.MenuItem, error) {
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
