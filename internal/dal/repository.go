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

type ListModels interface {
	[]models.MenuItem | []models.OrderItem | []models.InventoryItem
}

func list[T ListModels](filePath string) (T, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var items T

	err = json.Unmarshal(data, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func write[T ListModels](items T, filePath string) error {
	dat, err := json.MarshalIndent(items, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, dat, 0644)
	if err != nil {
		return err
	}

	return nil
}
