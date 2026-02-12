package models

import "errors"

type InventoryItem struct {
	IngredientID string  `json:"ingredient_id"`
	Name         string  `json:"name"`
	Quantity     float64 `json:"quantity"`
	Unit         string  `json:"unit"`
}

var (
	InventoryItemFound = errors.New("menu item not found")
)
