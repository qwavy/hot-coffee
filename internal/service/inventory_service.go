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
	if err != nil {
		return nil, err
	}

	return inventoryItems, nil
}

func (s *InventoryService) GetById(id string) (item *models.InventoryItem, err error) {
	inventoryItem, err := s.InventoryRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return inventoryItem, nil
}

func (s *InventoryService) Create(item models.InventoryItem) error {

	err := item.Validate()

	if err != nil {
		return err
	}

	err = s.InventoryRepository.Create(item)

	if err != nil {
		return err
	}

	return nil
}

func (s *InventoryService) Update(id string, item models.InventoryItem) error {
	err := item.Validate()

	if err != nil {
		return err
	}

	err = s.InventoryRepository.Update(id, item)

	fmt.Println(item)
	if err != nil {
		return err
	}

	return nil
}

func (s *InventoryService) DeleteById(id string) (err error) {
	return s.InventoryRepository.DeleteById(id)
}
