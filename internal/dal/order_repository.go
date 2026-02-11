package dal

import (
	"encoding/json"
	"os"

	"hot-coffee/models"
)

type OrderRepository struct {
	filePath string
}

func NewOrderRepository(filePath string) *OrderRepository {
	return &OrderRepository{filePath: filePath}
}

func (r *OrderRepository) write(orders []models.Order) error {
	data, err := json.MarshalIndent(orders, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(r.filePath, data, 0o644)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) listO() ([]models.Order, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var order []models.Order

	err = json.Unmarshal(data, &order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) GetAll() ([]models.Order, error) {
	return r.listO()
}

func (r *OrderRepository) GetById(productId string) (models.Order, error) {
	orderItems, err := r.listO()
	if err != nil {
		return models.Order{}, err
	}

	for _, orderItem := range orderItems {
		if orderItem.ID == productId {
			return orderItem, nil
		}
	}

	return models.Order{}, models.OrderItemNotFound
}

func (r *OrderRepository) DeleteById(productId string) error {
	orders, err := r.listO()
	if err != nil {
		return err
	}
	var newOrder []models.Order
	for _, order := range orders {
		if order.ID != productId {
			newOrder = append(newOrder, order)
		}
	}
	writerErr := r.write(newOrder)

	if writerErr != nil {
		return err
	}
	return nil
}
