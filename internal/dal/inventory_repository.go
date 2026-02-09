package dal

type InventoryRepository struct {
	filePath string
}

func NewInventoryRepository(filePath string) *InventoryRepository {
	return &InventoryRepository{filePath: filePath}
}
