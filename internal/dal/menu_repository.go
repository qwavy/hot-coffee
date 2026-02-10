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

func (r *MenuRepository) write(menuItems []models.MenuItem) error {
	dat, err := json.Marshal(menuItems)
	if err != nil {
		return err
	}
	err = os.WriteFile(r.filePath, dat, 0644)
	if err != nil {
		return err
	}

	return nil
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
			return menuItem, nil
		}
	}

	return models.MenuItem{}, models.MenuItemNotFound
}

func (r *MenuRepository) DeleteById(productId string) error {
	menuItems, err := r.list()

	if err != nil {
		return err
	}

	var newMenuItems []models.MenuItem

	for _, menuItem := range menuItems {
		if menuItem.ID != productId {
			newMenuItems = append(newMenuItems, menuItem)
		}
	}

	err = r.write(newMenuItems)

	if err != nil {
		return err
	}

	return nil
}
