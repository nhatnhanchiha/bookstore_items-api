package requests

import "github.com/nhatnhanchiha/bookstore_items-api/domain/items"

type UpdateItemRequest struct {
	Id                string
	Seller            int64
	Title             string            `json:"title"`
	Description       items.Description `json:"description"`
	Pictures          []items.Picture   `json:"pictures"`
	Video             string            `json:"video"`
	Price             float32           `json:"price"`
	AvailableQuantity int               `json:"available_quantity"`
	SoldQuantity      int               `json:"sold_quantity"`
	Status            string            `json:"status"`
}
