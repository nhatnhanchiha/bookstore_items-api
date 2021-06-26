package requests

import "github.com/nhatnhanchiha/bookstore_items-api/domain/items"

type PatchItemRequest struct {
	Id                string
	Seller            int64
	Title             *string         `json:"title,omitempty"`
	Description       *Description    `json:"description,omitempty"`
	Pictures          []items.Picture `json:"pictures,omitempty"`
	Video             *string         `json:"video,omitempty"`
	Price             *float32        `json:"price,omitempty"`
	AvailableQuantity *int            `json:"available_quantity,omitempty"`
	SoldQuantity      *int            `json:"sold_quantity,omitempty"`
	Status            *string         `json:"status,omitempty"`
}

type Description struct {
	PlainText *string `json:"plain_text,omitempty"`
	Html      *string `json:"html,omitempty"`
}
