package dal

import "hot-coffee/models"

type InventoryRepository struct {
	filePath string
}

func NewInventoryRepository(filePath string) *InventoryRepository {
	return &InventoryRepository{filePath: filePath}
}

func (r *InventoryRepository) GetAll() ([]models.InventoryItem, error) {
	return list[[]models.InventoryItem](r.filePath)
}
