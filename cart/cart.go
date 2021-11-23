package cart

import (
	"checkout_system/models"
	"checkout_system/upload"
	"fmt"
)

type ProductCart struct {
	Products []models.Product `json:"product"`
}

var AvailableProducts upload.Products

func New() *ProductCart {
	err := AvailableProducts.LoadProductPrices()
	if err != nil {
		fmt.Println("asd")
	}
	return new(ProductCart)
}

func (c *ProductCart) Contains(sku string) bool {
	for _, p := range c.Products {
		if p.SKU == sku {
			return true
		}
	}
	return false
}

func (c *ProductCart) Get(sku string) ProductCart {
	var data []models.Product
	for _, p := range c.Products {
		if p.SKU == sku {
			data = append(data, p)
		}
	}
	return ProductCart{Products: data}
}

func (c ProductCart) Len() int {
	return len(c.Products)
}

func (c *ProductCart) UpdateItemPrice(sku string, price float32) {
	for i, p := range c.Products {
		if p.SKU == sku {
			c.Products[i].Price = price
		}
	}
	return
}

func (c *ProductCart) UpdateOneItemPrice(sku string, price float32) {
	for i, p := range c.Products {
		if p.SKU == sku {
			if p.Price != 0 {
				c.Products[i].Price = price
				return
			} else {
				continue
			}
		}
	}
	return
}

func (c *ProductCart) Add(sku string) {
	for _, p := range AvailableProducts.Products {
		if p.SKU == sku {
			c.Products = append(c.Products, p)
		}
	}
	return
}

func (c *ProductCart) Total() float32 {
	var totalAmount float32
	for _, p := range c.Products {
		totalAmount += p.Price
	}
	return totalAmount
}
