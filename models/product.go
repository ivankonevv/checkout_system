package models

type Product struct {
	SKU   string  `json:"sku"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
