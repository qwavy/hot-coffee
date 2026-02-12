package service

import (
	"fmt"
	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type InventoryService struct {
	InventoryRepository *dal.InventoryRepository
}

func NewInventoryService(inventoryRepository *dal.InventoryRepository) *InventoryService {
	return &InventoryService{InventoryRepository: inventoryRepository}
}

func (s *InventoryService) GetAll() ([]models.InventoryItem, error) {
	inventoryItems, err := s.InventoryRepository.GetAll()
	fmt.Println(inventoryItems)
	if err != nil {
		return nil, err
	}

	return inventoryItems, nil
}
