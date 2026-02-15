package models

import "errors"

type Order struct {
	ID           string      `json:"order_id"`
	CustomerName string      `json:"customer_name"`
	Items        []OrderItem `json:"items"`
	Status       string      `json:"status"`
	CreatedAt    string      `json:"created_at"`
}

type OrderItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

var (
	OrderItemNotFound = errors.New("order item not found")
)

func (item *Order) Validate() error {
	if item.ID == "" {
		return errors.New("ID is required")
	}

	if item.CustomerName == "" {
		return errors.New("customer name is required")
	}

	if len(item.Items) == 0 {
		return errors.New("items is required")
	}

	if item.Status == "" {
		return errors.New("status is required")
	}

	if item.CreatedAt == "" {
		return errors.New("createdAt is required")
	}

	return nil
}
