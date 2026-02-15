package models

import "errors"

type MenuItem struct {
	ID          string               `json:"product_id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Price       float64              `json:"price"`
	Ingredients []MenuItemIngredient `json:"ingredients"`
}

type MenuItemIngredient struct {
	IngredientID string  `json:"ingredient_id"`
	Quantity     float64 `json:"quantity"`
}

var (
	MenuItemNotFound = errors.New("menu item not found")
)

func (item *MenuItem) Validate() error {
	if item.ID == "" {
		return errors.New("ID is required")
	}

	if item.Name == "" {
		return errors.New("name is required")
	}

	if item.Description == "" {
		return errors.New("Description is required")
	}

	if item.Price == 0 {
		return errors.New("unit is required")
	}

	if item.Price == 0 {
		return errors.New("unit is required")
	}

	if len(item.Ingredients) == 0 {
		return errors.New("ingredients is required")
	}

	return nil
}
