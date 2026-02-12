package dal

import (
	"hot-coffee/models"
)

type InventoryRepository struct {
	filePath string
}

func NewInventoryRepository(filePath string) *InventoryRepository {
	return &InventoryRepository{filePath: filePath}
}

func (r *InventoryRepository) GetAll() ([]models.InventoryItem, error) {
	return list[[]models.InventoryItem](r.filePath)
}

func (r *InventoryRepository) GetById(id string) (*models.InventoryItem, error) {
	inventoryItems, err := list[[]models.InventoryItem](r.filePath)

	if err != nil {
		return nil, err
	}

	for _, inventoryItem := range inventoryItems {
		if inventoryItem.IngredientID == id {
			return &inventoryItem, nil
		}
	}

	return nil, models.InventoryItemFound
}

func (r *InventoryRepository) DeleteById(id string) error {
	inventoryItems, err := list[[]models.InventoryItem](r.filePath)

	if err != nil {
		return err
	}

	var newInventoryItems []models.InventoryItem

	for _, inventoryItem := range inventoryItems {
		if inventoryItem.IngredientID != id {
			newInventoryItems = append(newInventoryItems, inventoryItem)
		}
	}

	err = write(newInventoryItems, r.filePath)

	if err != nil {
		return err
	}

	return nil
}
