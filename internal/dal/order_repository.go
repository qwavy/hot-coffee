package dal

import (
	"encoding/json"
	"hot-coffee/models"
	"os"
)

type OrderRepository struct {
	filePath string
}

func NewOrderRepository(filePath string) *OrderRepository {
	return &OrderRepository{filePath: filePath}
}
func (r *OrderRepository) listO() ([]models.OrderItem, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var orderItem []models.OrderItem

	err = json.Unmarshal(data, &orderItem)
	if err != nil {
		return nil, err
	}

	return orderItem, nil
}
func (r *OrderRepository) GetAll() ([]models.OrderItem, error) {
	return r.listO()
}