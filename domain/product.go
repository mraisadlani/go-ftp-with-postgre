package domain

import "time"

type Product struct {
	ID uint `json:"product_id"`
	ProductCode string `json:"product_code"`
	ProductName string `json:"product_name"`
	ProductSlug string `json:"product_slug"`
	ProductDescription string `json:"product_description"`
	Qty uint `json:"qty"`
	MinQty uint `json:"min_qty"`
	MaxQty uint `json:"max_qty"`
	Weight uint `json:"weight"`
	Volume *uint `json:"volume"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt *time.Time `json:"update_at"`
}

type ProductDTO struct {
	ProductCode string `json:"product_code"`
	ProductName string `json:"product_name"`
	ProductSlug string `json:"product_slug"`
	ProductDescription string `json:"product_description"`
	Qty uint `json:"qty"`
	MinQty uint `json:"min_qty"`
	MaxQty uint `json:"max_qty"`
	Weight uint `json:"weight"`
	Volume uint `json:"volume"`
}