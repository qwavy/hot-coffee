package dal

import (
	"hot-coffee/models"
)

type MenuRepository struct {
	filePath string
}

func NewMenuRepository(filePath string) *MenuRepository {
	return &MenuRepository{filePath: filePath}
}

func (r *MenuRepository) GetAll() ([]models.MenuItem, error) {
	return list(r.filePath)
}

func (r *MenuRepository) GetById(productId string) (*models.MenuItem, error) {
	menuItems, err := list(r.filePath)

	if err != nil {
		return nil, err
	}

	for _, menuItem := range menuItems {
		if menuItem.ID == productId {
			return &menuItem, nil
		}
	}

	return nil, models.MenuItemNotFound
}

func (r *MenuRepository) DeleteById(productId string) error {
	menuItems, err := list(r.filePath)

	if err != nil {
		return err
	}

	var newMenuItems []models.MenuItem

	for _, menuItem := range menuItems {
		if menuItem.ID != productId {
			newMenuItems = append(newMenuItems, menuItem)
		}
	}

	err = write(newMenuItems, r.filePath)

	if err != nil {
		return err
	}

	return nil
}

func (r *MenuRepository) CreateItem(menuItem models.MenuItem) error {
	menuItems, err := list(r.filePath)

	if err != nil {
		return err
	}

	menuItems = append(menuItems, menuItem)
	err = write(menuItems, r.filePath)

	if err != nil {
		return err
	}

	return nil
}

func (r *MenuRepository) UpdateItem(productId string, newMenuItem models.MenuItem) error {
	menuItems, err := list(r.filePath)

	if err != nil {
		return err
	}

	for i, menuItem := range menuItems {
		if menuItem.ID == productId {
			newMenuItem.ID = productId
			menuItems[i] = newMenuItem

			err = write(menuItems, r.filePath)

			if err != nil {
				return err
			}

			return nil
		}
	}

	return models.MenuItemNotFound

}
