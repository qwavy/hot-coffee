package models

import "errors"

type InventoryItem struct {
	IngredientID string  `json:"ingredient_id"`
	Name         string  `json:"name"`
	Quantity     float64 `json:"quantity"`
	Unit         string  `json:"unit"`
}

var (
	InventoryItemFound         = errors.New("inventory item not found")
	InventoryItemAlreadyExists = errors.New("inventory item with this id already exists")
)

func (item *InventoryItem) Validate() error {
	if item.IngredientID == "" {
		return errors.New("IngredientID is required")
	}

	if item.Name == "" {
		return errors.New("name is required")
	}

	if item.Quantity == 0 {
		return errors.New("quantity is required")
	}

	if item.Quantity == 0 {
		return errors.New("unit is required")
	}

	return nil
}
