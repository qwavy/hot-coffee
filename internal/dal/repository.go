package dal

import (
	"encoding/json"
	"hot-coffee/models"
	"os"
)

type Repositories struct {
	Inventory *InventoryRepository
	Menu      *MenuRepository
	Order     *OrderRepository
}

func NewRepositories(inventoryRepositoryFilepath, menuRepositoryFilepath, orderRepositoryFilepath string) *Repositories {
	return &Repositories{
		Inventory: NewInventoryRepository(inventoryRepositoryFilepath),
		Menu:      NewMenuRepository(menuRepositoryFilepath),
		Order:     NewOrderRepository(orderRepositoryFilepath),
	}
}

func list(filePath string) ([]models.MenuItem, error) {
	data, err := os.ReadFile(filePath)
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

func write(menuItems []models.MenuItem, filePath string) error {
	dat, err := json.MarshalIndent(menuItems, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, dat, 0644)
	if err != nil {
		return err
	}

	return nil
}
