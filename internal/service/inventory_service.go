package service

import "hot-coffee/internal/dal"

type InventoryService struct {
	InventoryRepository dal.InventoryRepository
}

func NewInventoryService(inventoryRepository dal.InventoryRepository) *InventoryService {
	return &InventoryService{InventoryRepository: inventoryRepository}
}
